import java.util.ArrayList;
import java.util.Collections;
import java.util.Comparator;
import java.util.List;

public class Day1 implements Day{

    public int part1(String input) {
        List<Integer> caloriesByElf = getCaloriesByElf(input);
        return caloriesByElf.stream()
                .max(Comparator.naturalOrder())
                .get();
    }

    public int part2(String input) {
        List<Integer> caloriesByElf = getCaloriesByElf(input);
        return caloriesByElf.stream()
                .sorted(Collections.reverseOrder())
                .limit(3)
                .reduce(0, Integer::sum);
    }

    private List<Integer> getCaloriesByElf(String listOfCalories) {
        List<Integer> caloriesByElf = new ArrayList<>();

        for (String singleElfCalories : listOfCalories.split("\n\n")) {
            int sum = 0;
            for (String calories : singleElfCalories.split("\n")) {
                sum += Integer.parseInt(calories);
            }
            caloriesByElf.add(sum);
        }

        return caloriesByElf;
    }
}
