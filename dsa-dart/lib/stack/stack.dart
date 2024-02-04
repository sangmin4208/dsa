class Stack<T> {
  final List<T> _stack = [];

  void push(T value) {
    _stack.add(value);
  }

  T pop() {
    return _stack.removeLast();
  }

  T peek() {
    return _stack.last;
  }

  bool get isEmpty => _stack.isEmpty;

  int get size => _stack.length;
}
