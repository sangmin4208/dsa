import { describe, expect, test } from "vitest";

import { Stack } from "./stack";

describe("Stack", () => {
  test("should create an empty stack", () => {
    const stack = new Stack();
    expect(stack.size()).toBe(0);
    expect(stack.isEmpty()).toBe(true);
  });

  test("should push elements to the stack", () => {
    const stack = new Stack();
    stack.push(1);
    stack.push(2);
    expect(stack.size()).toBe(2);
    expect(stack.isEmpty()).toBe(false);
  });

  test("should pop elements from the stack", () => {
    const stack = new Stack();
    stack.push(1);
    stack.push(2);
    expect(stack.pop()).toBe(2);
    expect(stack.pop()).toBe(1);
    expect(stack.size()).toBe(0);
  });

  test("should peek the top element of the stack", () => {
    const stack = new Stack();
    stack.push(1);
    stack.push(2);
    expect(stack.peek()).toBe(2);
    stack.pop();
    expect(stack.peek()).toBe(1);
  });
});
