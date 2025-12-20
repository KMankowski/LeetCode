package solutions.p252;

import java.util.List;

class Solution {
    static class Interval {
        int start;
        int end;

        Interval(int start, int end) {
            this.start = start;
            this.end = end;
        }
    }

    boolean canAttendMeetings(List<Interval> intervals) {
        intervals.sort((i, j) -> Integer.compare(i.start, j.start));

        for (int i = 0; i < intervals.size() - 1; i++) {
            if (intervals.get(i).end > intervals.get(i + 1).start) {
                return false;
            }
        }

        return true;
    }
}