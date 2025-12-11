package solutions.p215;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.stream.Stream;

import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

public class SolutionTest {
    static Stream<Arguments> getTestCases() {
        return Stream.of(
                Arguments.of(new int[] { 3, 2, 1, 5, 6, 4 }, 2, 5),
                Arguments.of(new int[] { 3, 2, 3, 1, 2, 4, 5, 5, 6 }, 4, 4));
    }

    @ParameterizedTest
    @MethodSource("getTestCases")
    void test(int inpNums[], int inpK, int expKthLargest) {
        assertEquals(expKthLargest, new Solution().findKthLargest(inpNums, inpK));
    }
}
