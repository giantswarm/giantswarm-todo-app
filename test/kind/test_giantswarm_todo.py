import time
import pytest
from typing import List
from kubetest.client import TestClient
from kubetest.objects import Service, Deployment


todo_timeout: int = 90


@pytest.fixture(scope="function")
def services(kube: TestClient) -> List[Service]:
    services = kube.get_services(namespace="default")
    for service_name in ["apiserver", "giantswarm-todo-app-mysql", "todomanager"]:
        service = services.get(service_name)
        assert service is not None
        service.wait_until_ready(timeout=todo_timeout)
    return services


@pytest.fixture(scope="function")
def deployments(kube: TestClient) -> List[Deployment]:
    deployments = kube.get_deployments(namespace="default")
    for deploy_name in ["apiserver", "giantswarm-todo-app-mysql", "todomanager"]:
        deployment = deployments.get(deploy_name)
        assert deployment is not None
        deployment.wait_until_ready(timeout=todo_timeout)
    return deployments


def test_deployments(deployments: List[Deployment]):
    for deploy_name in ["apiserver", "giantswarm-todo-app-mysql", "todomanager"]:
        deployment = deployments.get(deploy_name)
        assert deployment.is_ready()


def test_services(services: List[Service]):
    for service_name in ["apiserver", "giantswarm-todo-app-mysql", "todomanager"]:
        service = services.get(service_name)
        assert service.is_ready()


# By injecting fixtures, we can be sure that all deployments and the service are "Ready"
@pytest.mark.flaky(reruns=5, reruns_delay=2)
def test_get_todos(services: List[Service], deployments: List[Deployment]):
    # unfortunately, when services and deployments are ready, traffic forwarding doesn't yet
    # work fo 100% :( That's why we need a retry.
    data, status, headers = services.get("apiserver").proxy_http_get("v1/todo")
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
