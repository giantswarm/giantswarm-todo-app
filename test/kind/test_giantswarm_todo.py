import pytest
from typing import List
from kubetest.client import TestClient
from kubetest.objects import Service, Deployment


todo_timeout: int = 90


@pytest.fixture(scope="function")
def apiserver_service(kube: TestClient) -> Service:
    service = kube.get_services(namespace="default").get("apiserver")
    assert service is not None
    service.wait_until_ready(timeout=todo_timeout)
    return service


@pytest.fixture(scope="function")
def deployments(kube: TestClient) -> List[Deployment]:
    deployments = kube.get_deployments(namespace="default")
    for deploy_name in ["apiserver", "giantswarm-todo-app-mysql", "todomanager"]:
        deployment = deployments.get(deploy_name)
        deployment.wait_until_ready(timeout=todo_timeout)
    return deployments

def test_deployments(deployments: List[Deployment]):
    for deploy_name in ["apiserver", "giantswarm-todo-app-mysql", "todomanager"]:
        deployment = deployments.get(deploy_name)
        assert deployment.is_ready()


def test_apiserver_service(apiserver_service: Service):
    assert apiserver_service.is_ready()


# By injecting fixtures, we can be sure that all deployments and the service are "Ready"
def test_get_todos(apiserver_service: Service, deployments: List[Deployment]):
    data, status, headers = apiserver_service.proxy_http_get("v1/todo")
    assert data == None
    assert status == 200
    assert 'Content-Type' in headers
    assert headers['Content-Type'] == 'application/json; charset=utf-8'


# def test_create_todo_entry(apiserver_service: Service):
#    header_params = {}
#    header_params['Content-Type'] = 'application/json'
#    body = '{"Text":"testing"}'
#    apiserver_service.proxy_http_post(
#        "v1/todo", header_params=header_params, body=body)
#
