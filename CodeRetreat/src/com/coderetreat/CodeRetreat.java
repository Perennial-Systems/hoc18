package com.coderetreat;

import java.util.Scanner;

public class CodeRetreat {

	public static void main(String[] args) {

		Scanner scanner = new Scanner(System.in);
		// Users values
		System.out.println("####### Please Enter the start value from where you want to get the prime No ########");
		int x = Integer.parseInt(scanner.next());
		System.out.println("####### Please Enter the end value till you want to get the prime No ########");
		int y = Integer.parseInt(scanner.next());
		
		// Check validations
		if(x <= 1 || y <= 1 ) {
			System.out.println("Please enter number greator than zero");
		}else if(x >= y) {
			System.out.println("First value must be smaller than second value");
		}else {
			incrementNumber(x,y);
		}
		scanner.close();
	}
	
	
	/** This method will increment the value and check the prime number */
	static void incrementNumber(int startNumber, int endNumber) {
		if(startNumber <= endNumber) {
			checkPrime(startNumber,1);
			startNumber++;
			incrementNumber(startNumber, endNumber);
		}
	}
	
	/** This method will check for the prime no*/
	static boolean checkPrime(int number, int decrementValue) {
		int num = number - decrementValue;
		if(num == 1) {
			System.out.println("This is the prime number : " + number);
			return true;
		}
		
		if(number%num == 0) {
			System.out.println("This is not the prime number : " + number);
			return false;
		}else {
			checkPrime(number,++decrementValue);
		}
		return false;
	}

}
