package solutions.p252;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.Arrays;
import java.util.List;
import java.util.stream.Stream;

import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import solutions.p252.Solution.Interval;

public class SolutionTest {
    static Stream<Arguments> getTestCases() {
        return Stream.of(
                Arguments.of(Arrays.asList(new Interval(0, 30), new Interval(5, 10), new Interval(15, 20)),
                        false),
                Arguments.of(Arrays.asList(new Interval(5, 8), new Interval(9, 15)),
                        true));
    }

    @ParameterizedTest
    @MethodSource("getTestCases")
    void test(List<Interval> intervals, boolean canAttend) {
        assertEquals(canAttend, new Solution().canAttendMeetings(intervals));
    }
}
