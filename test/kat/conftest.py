import pytest
import yaml

from typing import Any


def pytest_addoption(parser):
    parser.addoption("--values-file", action="store")
    parser.addoption("--chart-name", action="store")


@pytest.fixture(scope="module")
def chart_name(pytestconfig) -> str:
    return pytestconfig.getoption("chart-name")


@pytest.fixture(scope="module")
def values_file_path(pytestconfig) -> str:
    return pytestconfig.getoption("values-file")


@pytest.fixture(scope="module")
def values_file(values_file_path) -> Any:
    with open(values_file_path) as f:
        return yaml.safe_load(f)
