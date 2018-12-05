package com.hourofcode;

import java.util.ArrayList;
import java.util.List;

public class PrimeNumber {

	public static List<Integer> primeList = new ArrayList<>();

	public static void main(String[] args) {
		int firstNo = 1;
		int lastNo = 100;
		printPrimeNumber(firstNo, lastNo);
		System.out.println(primeList);
	}

	private static void printPrimeNumber(int firstNo, int lastNo) {
		prime(firstNo, 2, lastNo);
	}

	static void prime(int sampleInput, int i, int lastNo) {
		if (sampleInput <= lastNo) {
			if (i < sampleInput) {
				if (sampleInput % i != 0) {
					prime(sampleInput, ++i, lastNo);
				} else {

				}
			}
			primeList.add(sampleInput);
			++sampleInput;
			prime(sampleInput, 2, lastNo);
		}
	}
}
