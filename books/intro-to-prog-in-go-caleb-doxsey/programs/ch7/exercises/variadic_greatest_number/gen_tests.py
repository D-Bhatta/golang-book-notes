"""Generate tests for this module."""

import random
from typing import TypedDict
import json

# Limit of the largest numbered test case. LImited to int32 in go.
LARGEST_NUM = 100_000
SMALLEST_NUM = -10_000
FILE_NAME = "test_cases.json"
NUM_TEST_CASES = 100


class TestCase(TypedDict):
    """Contain data for each test case."""

    id: int
    largest: int
    numbers: list[int]
    err_expected: str | None


def create_test_case(index: int) -> TestCase:
    """Create a TestCase with upto 100 numbers."""
    big = random.randint(SMALLEST_NUM, LARGEST_NUM)
    length = random.randint(1, 100)
    position = random.randint(0, length)
    if position == length:
        position = length - 1

    nums = [random.randint(SMALLEST_NUM, big) for _ in range(length)]
    nums[position] = big

    test_case: TestCase = TestCase(
        id=index, largest=big, numbers=nums, err_expected=None
    )
    return test_case


tests: dict[str, list[TestCase]] = {
    "tests": [create_test_case(index=i) for i in range(1, NUM_TEST_CASES + 1)]
}


tests_json: str = json.dumps(tests)

with open(FILE_NAME, "w", encoding="utf-8") as fh:
    fh.write(tests_json)
