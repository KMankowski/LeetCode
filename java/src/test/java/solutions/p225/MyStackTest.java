package solutions.p225;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;

import org.junit.jupiter.api.Test;

public class MyStackTest {

    @Test
    void test() {
        MyStack stack = new MyStack();
        stack.push(1);
        stack.push(2);
        assertEquals(2, stack.top());
        assertEquals(2, stack.pop());
        assertFalse(stack.empty());
    }

}
