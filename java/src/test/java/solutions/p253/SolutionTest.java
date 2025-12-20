package solutions.p253;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.List;
import java.util.stream.Stream;

import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import solutions.p253.Solution.Interval;

public class SolutionTest {
    static Stream<Arguments> getTestCases() {
        return Stream.of(
                Arguments.of(
                        List.of(
                                new Interval(0, 40),
                                new Interval(5, 10),
                                new Interval(15, 20)),
                        2),
                Arguments.of(
                        List.of(
                                new Interval(4, 9)),
                        1));
    }

    @ParameterizedTest
    @MethodSource("getTestCases")
    void test(List<Interval> intervals, int expDaysNeeded) {
        assertEquals(expDaysNeeded, new Solution().minMeetingRooms(intervals));
    }
}
