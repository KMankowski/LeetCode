package solutions.p225;

import java.util.ArrayDeque;
import java.util.Deque;

public class MyStack {

    private final Deque<Integer> firstQueue = new ArrayDeque<>();
    private final Deque<Integer> secondQueue = new ArrayDeque<>();
    private boolean isFirstQueueActive = true;

    public void push(int x) {
        if (isFirstQueueActive) {
            firstQueue.add(x);
        } else {
            secondQueue.add(x);
        }
    }

    public int pop() {
        int pop;
        if (isFirstQueueActive) {
            pop = pop(firstQueue, secondQueue);
        } else {
            pop = pop(secondQueue, firstQueue);
        }
        isFirstQueueActive = !isFirstQueueActive;
        return pop;
    }

    private int pop(Deque<Integer> currQueue, Deque<Integer> newQueue) {
        if (newQueue.size() != 0) {
            throw new RuntimeException();
        }
        while (currQueue.size() > 1) {
            newQueue.add(currQueue.remove());
        }
        return currQueue.remove();
    }

    public int top() {
        int top;
        if (isFirstQueueActive) {
            top = top(firstQueue, secondQueue);
        } else {
            top = top(secondQueue, firstQueue);
        }
        isFirstQueueActive = !isFirstQueueActive;
        return top;
    }

    private int top(Deque<Integer> currQueue, Deque<Integer> newQueue) {
        if (newQueue.size() != 0) {
            throw new RuntimeException();
        }
        while (currQueue.size() > 1) {
            newQueue.add(currQueue.remove());
        }
        int top = currQueue.remove();
        newQueue.add(top);
        return top;
    }

    public boolean empty() {
        return isFirstQueueActive ? (firstQueue.size() == 0) : (secondQueue.size() == 0);
    }
}
