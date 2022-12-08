import org.junit.Test;
import java.io.IOException;
import java.net.URISyntaxException;

import static org.junit.Assert.assertEquals;

public class Day1Test {

    private final Day1 day1 = new Day1();

    @Test
    public void test_part1_example() throws IOException, URISyntaxException {
        String input = new String(this.getClass().getResourceAsStream("Day1_example").readAllBytes());
        assertEquals(day1.part1(input), 24000);
    }

    @Test
    public void test_part1() throws IOException {
        String input = new String(this.getClass().getResourceAsStream("Day1_part1").readAllBytes());
        assertEquals(day1.part1(input), 69883);
    }

    @Test
    public void test_part2_example() throws IOException {
        String input = new String(this.getClass().getResourceAsStream("Day1_example").readAllBytes());
        assertEquals(day1.part2(input), 45000);
    }

    @Test
    public void test_part2() throws IOException {
        String input = new String(this.getClass().getResourceAsStream("Day1_part2").readAllBytes());
        assertEquals(day1.part2(input), 207576);
    }

}
