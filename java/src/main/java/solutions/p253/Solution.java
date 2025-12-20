package solutions.p253;

import java.util.ArrayList;
import java.util.List;

public class Solution {
    public static class Interval {
        int start;
        int end;

        Interval(int start, int end) {
            this.start = start;
            this.end = end;
        }

        public String toString() {
            return start + " " + end;
        }
    }

    public int minMeetingRooms(List<Interval> intervals) {
        enum EventType {
            START,
            END,
        }

        record Event(EventType type, int position) {
        }

        List<Event> events = new ArrayList<>(intervals.size() * 2);
        intervals.forEach(interval -> {
            events.add(new Event(EventType.START, interval.start));
            events.add(new Event(EventType.END, interval.end));
        });
        events.sort((i, j) -> {
            int cmp = Integer.compare(i.position, j.position);
            if (cmp != 0 || i.type != j.type) {
                return cmp;
            }
            if (i.type == EventType.START) {
                return -1;
            } else {
                return 1;
            }
        });

        int maxDaysNeeded = 0;
        int currDaysNeeded = 0;
        for (Event event : events) {
            switch (event.type) {
                case START -> {
                    currDaysNeeded++;
                    if (currDaysNeeded > maxDaysNeeded) {
                        maxDaysNeeded = currDaysNeeded;
                    }
                }
                case END -> {
                    currDaysNeeded--;
                }
            }
        }

        return maxDaysNeeded;
    }
}
