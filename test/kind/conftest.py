import pytest


def pytest_addoption(parser):
    parser.addoption("--values-file", action="store")

@pytest.fixture(scope="module")
def values_file_path(pytestconfig):
    return pytestconfig.getoption("values-file")
