import json
import time
from typing import List

import pytest
from pykube import Service, Deployment, HTTPClient
from pytest_helm_charts.clusters import Cluster
from requests import Response

todo_timeout: int = 90


@pytest.fixture(scope="function")
def deployments(kube_cluster: Cluster) -> List[Deployment]:
    retries = 0
    all_ready = False
    my_deployments: List[Deployment] = []
    while retries < todo_timeout:
        deployments_response = Deployment.objects(kube_cluster.kube_client).filter(
            namespace="default"
        )
        retries += 1
        my_deployments = []
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
    services_response = Service.objects(kube_cluster.kube_client).filter(
        namespace="default"
    )
    for service_name in ["apiserver", "giantswarm-todo-app-mysql", "todomanager"]:
        service = services_response.get_by_name(service_name)
        assert service is not None


def test_deployments(deployments: List[Deployment]):
    for d in deployments:
        assert d.obj["status"]["availableReplicas"] > 0


# By injecting fixtures, we can be sure that all deployments and the service are "Ready"
# @pytest.mark.flaky(reruns=10, reruns_delay=3)
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


def _proxy_http_request(
    client: HTTPClient, srv: Service, method, path, **kwargs
) -> Response:
    """Template request to proxy of a Service.
    Args:
        :param client: HTTPClient to use.
        :param srv: Service you want to proxy.
        :param method: The http request method e.g. 'GET', 'POST' etc.
        :param path: The URI path for the request.
        :param kwargs: Keyword arguments for the proxy_http_get function.
    Returns:
        The Response data.
    """
    if "port" in kwargs:
        port = kwargs["port"]
    else:
        port = srv.obj["spec"]["ports"][0]["port"]
    kwargs["url"] = f"services/{srv.name}:{port}/proxy/{path}"
    kwargs["namespace"] = srv.namespace
    kwargs["version"] = srv.version
    return client.request(method, **kwargs)


def proxy_http_get(client: HTTPClient, srv: Service, path: str, **kwargs) -> Response:
    """Issue a GET request to proxy of a Service.
    Args:
        :param client: HTTPClient to use.
        :param srv: Service you want to proxy.
        :param path: The URI path for the request.
        :param kwargs: Keyword arguments for the proxy_http_get function.
    Returns:
        The response data.
    """
    return _proxy_http_request(client, srv, "GET", path, **kwargs)


def proxy_http_post(client: HTTPClient, srv: Service, path: str, **kwargs) -> Response:
    """Issue a POST request to proxy of a Service.
    Args:
        :param client: HTTPClient to use.
        :param srv: Service you want to proxy.
        :param path: The URI path for the request.
        :param kwargs: Keyword arguments for the proxy_http_get function.
    Returns:
        The response data.
    """
    return _proxy_http_request(client, srv, "POST", path, **kwargs)


def proxy_http_put(client: HTTPClient, srv: Service, path: str, **kwargs) -> Response:
    """Issue a PUT request to proxy of a Service.
    Args:
        :param client: HTTPClient to use.
        :param srv: Service you want to proxy.
        :param path: The URI path for the request.
        :param kwargs: Keyword arguments for the proxy_http_get function.
    Returns:
        The response data.
    """
    return _proxy_http_request(client, srv, "PUT", path, **kwargs)


def proxy_http_delete(
    client: HTTPClient, srv: Service, path: str, **kwargs
) -> Response:
    """Issue a DELETE request to proxy of a Service.
    Args:
        :param client: HTTPClient to use.
        :param srv: Service you want to proxy.
        :param path: The URI path for the request.
        :param kwargs: Keyword arguments for the proxy_http_get function.
    Returns:
        The response data.
    """
    return _proxy_http_request(client, srv, "DELETE", path, **kwargs)
