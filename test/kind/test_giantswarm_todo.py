import pytest
from kubetest.client import TestClient
from kubetest.objects import Service


@pytest.fixture(scope="function")
def apiserver_service(kube: TestClient) -> Service:
    return kube.get_services(namespace="default").get("apiserver")


def test_deploy(kube: TestClient):
    deployments = kube.get_deployments(namespace="default")
    for deploy_name in ["apiserver", "giantswarm-todo-app-mysql", "todomanager"]:
        deployment = deployments.get(deploy_name)
        kube.wait_until_created(deployment, timeout=60)
        assert deployment.is_ready()


def test_apiserver_service(kube: TestClient):
    service = kube.get_services(namespace="default").get("apiserver")
    assert service is not None
    assert service.is_ready()


def test_get_todos(apiserver_service: Service):
    data, status, headers = apiserver_service.proxy_http_get("v1/todo")
    assert data == None
    assert status == 200
    assert 'Content-Type' in headers
    assert headers['Content-Type'] == 'application/json; charset=utf-8'


#def test_create_todo_entry(apiserver_service: Service):
#    header_params = {}
#    header_params['Content-Type'] = 'application/json'
#    body = '{"Text":"testing"}'
#    apiserver_service.proxy_http_post(
#        "v1/todo", header_params=header_params, body=body)
#