import time
from typing import List

import pytest
from pykube import Service, Deployment
from pytest_helm_charts.clusters import Cluster

todo_timeout: int = 90


@pytest.fixture(scope="function")
def deployments(kube_cluster: Cluster) -> List[Deployment]:
    retries = 0
    all_ready = False
    while retries < todo_timeout:
        deployments_response = Deployment.objects(kube_cluster.kube_client).filter(namespace="default")
        retries += 1
        my_deployments: List[Deployment] = []
        for deploy_name in ["apiserver", "giantswarm-todo-app-mysql", "todomanager"]:
            deployment = deployments_response.get_by_name(deploy_name)
            assert deployment is not None
            my_deployments.append(deployment)
        all_ready = all(d.ready for d in my_deployments)
        if all_ready:
            break
        time.sleep(1)

    if not all_ready:
        raise TimeoutError("Error waiting for deployments to become 'ready'.")

    return my_deployments


def test_services(kube_cluster: Cluster):
    services_response = Service.objects(kube_cluster.kube_client).filter(namespace="default")
    for service_name in ["apiserver", "giantswarm-todo-app-mysql", "todomanager"]:
        service = services_response.get_by_name(service_name)
        assert service is not None


def test_deployments(deployments: List[Deployment]):
    for deploy_name in ["apiserver", "giantswarm-todo-app-mysql", "todomanager"]:
        deployment = deployments.get(deploy_name)
        assert deployment.is_ready()

## By injecting fixtures, we can be sure that all deployments and the service are "Ready"
# @pytest.mark.flaky(reruns=10, reruns_delay=3)
# def test_get_todos(services: List[Service], deployments: List[Deployment]):
#    # unfortunately, when services and deployments are ready, traffic forwarding doesn't yet
#    # work fo 100% :( That's why we need a retry.
#    data, status, headers = services.get("apiserver").proxy_http_get("v1/todo")
#    assert data == None
#    assert status == 200
#    assert 'Content-Type' in headers
#    assert headers['Content-Type'] == 'application/json; charset=utf-8'


# def test_create_todo_entry(apiserver_service: Service):
#    header_params = {}
#    header_params['Content-Type'] = 'application/json'
#    body = '{"Text":"testing"}'
#    apiserver_service.proxy_http_post(
#        "v1/todo", header_params=header_params, body=body)
#
