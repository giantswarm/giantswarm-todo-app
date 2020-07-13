import json
from typing import List

import pytest
from pykube import Service, Deployment
from pytest_helm_charts.clusters import Cluster
from pytest_helm_charts.utils import wait_for_deployments_to_run, proxy_http_get, proxy_http_post, proxy_http_delete
from requests import Response

todo_timeout: int = 90


@pytest.fixture(scope="function")
def deployments(kube_cluster: Cluster) -> List[Deployment]:
    return wait_for_deployments_to_run(kube_cluster.kube_client,
                                       ["apiserver", "giantswarm-todo-app-mysql", "todomanager"],
                                       "default",
                                       todo_timeout)


def test_services(kube_cluster: Cluster):
    services_response = Service.objects(kube_cluster.kube_client).filter(
        namespace="default"
    )
    for service_name in ["apiserver", "giantswarm-todo-app-mysql", "todomanager"]:
        service = services_response.get_by_name(service_name)
        assert service is not None


def test_deployments(deployments: List[Deployment]):
    for d in deployments:
        assert int(d.obj["status"]["availableReplicas"]) > 0


# By injecting fixtures, we can be sure that all deployments and the service are "Ready"
@pytest.mark.flaky(reruns=10, reruns_delay=3)
@pytest.mark.usefixtures("deployments")
def test_get_todos(kube_cluster: Cluster):
    # unfortunately, when services and deployments are ready, traffic forwarding doesn't yet
    # work fo 100% :( That's why we need a retry.
    apiserver_service = (
        Service.objects(kube_cluster.kube_client)
            .filter(namespace="default")
            .get(name="apiserver")
    )
    res = proxy_http_get(kube_cluster.kube_client, apiserver_service, "v1/todo")
    assert res is not None
    assert res.content == b"null\n"
    assert res.status_code == 200
    assert "Content-Type" in res.headers
    assert res.headers["Content-Type"] == "application/json; charset=utf-8"


@pytest.mark.flaky(reruns=10, reruns_delay=3)
@pytest.mark.usefixtures("deployments")
def test_create_delete_todo_entry(kube_cluster: Cluster):
    apiserver_service = (
        Service.objects(kube_cluster.kube_client)
            .filter(namespace="default")
            .get(name="apiserver")
    )
    body = '{"Text":"testing"}'
    headers = {"Content-Type": "application/json"}
    res = proxy_http_post(
        kube_cluster.kube_client,
        apiserver_service,
        "v1/todo",
        data=body,
        headers=headers,
    )
    assert res is not None
    assert res.status_code == 200
    todo_id = json.loads(res.text)["id"]
    res = proxy_http_delete(
        kube_cluster.kube_client, apiserver_service, f"v1/todo/{todo_id}"
    )
    assert res is not None
    assert res.status_code == 200
