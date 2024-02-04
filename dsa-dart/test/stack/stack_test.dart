import 'package:dsa_dart/dsa_dart.dart';
import 'package:dsa_dart/stack/stack.dart';
import 'package:test/test.dart';

void main() {
  test('calculate', () {
    expect(calculate(), 42);
  });

  test('Stack - Push and Pop', () {
    Stack<int> stack = Stack<int>();
    stack.push(1);
    stack.push(2);
    stack.push(3);

    expect(stack.pop(), 3);
    expect(stack.pop(), 2);
    expect(stack.pop(), 1);
  });

  test('Stack - Peek', () {
    Stack<int> stack = Stack<int>();
    stack.push(1);
    stack.push(2);
    stack.push(3);

    expect(stack.peek(), 3);
    expect(stack.peek(), 3);
  });

  test('Stack - isEmpty', () {
    Stack<int> stack = Stack<int>();

    expect(stack.isEmpty, true);

    stack.push(1);

    expect(stack.isEmpty, false);
  });

  test('Stack - size', () {
    Stack<int> stack = Stack<int>();

    expect(stack.size, 0);

    stack.push(1);
    stack.push(2);
    stack.push(3);

    expect(stack.size, 3);
  });
}
