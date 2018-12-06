import java.io.DataInputStream;
import java.io.IOException;

public class Problem1 {

	private void printPrime(int start, int end) {
		if (start < end) {
			checkNumber(start);
			start++;
			printPrime(start, end);
		}
	}

	private void checkNumber(int num) {
		int init = 2;
		checkPrime(init, num);
	}

	private void checkPrime(int init, int num) {
		if (init < num) {
			if ((num % init) == 0) {
				return;
			} else {
				init++;
				this.checkPrime(init, num);
			}
		} else {
			System.out.println(num + " is prime number");
		}
	}

	public static void main(String[] args) {

		DataInputStream dis = new DataInputStream(System.in);

		int start = 10;
		int end = 20;

		try {
			System.out.println("Enter Starting number:");
			start = Integer.parseInt(dis.readLine());

			System.out.println("Enter last number:");
			end = Integer.parseInt(dis.readLine());

		} catch (IOException e) {
			e.printStackTrace();
		}

		Problem1 p = new Problem1();
		p.printPrime(start, end);

	}
}
