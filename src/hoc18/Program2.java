package hoc18;

public class Program2 {
	public static void main(String[] args) {
		// take input
		String input = "aWk60#ase2> o_0^4alis5L";
		String[] inputArr = input.split(" ");

		for (int i = 0; i < inputArr.length; i++) {
			String element = inputArr[i];
			char[] charArr = element.toCharArray();
			int value = 0;
			for (int j = 0; j < charArr.length; j++) {
				int asciNumber = charArr[j];
				if (asciNumber >= 48 && asciNumber <= 55) {
					System.out.println();
					char number = charArr[j];
					value = value + number;
				}
			}
			value = value - 48;
			System.out.print(Character.toString((char) value));
		}

	}
}
