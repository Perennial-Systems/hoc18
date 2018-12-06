package com.hourofcode;

import java.util.ArrayList;
import java.util.List;

public class PrimeNumber {

	public static List<Integer> primeList = new ArrayList<>();

	public static void main(String[] args) {
		int firstNo = 1;
		int lastNo = 100;
		isPrime(firstNo, 2, lastNo);
		System.out.println(primeList);
	}

	static void isPrime(int sampleInput, int i, int lastNo) {
		if (sampleInput <= lastNo) {
			if (i < sampleInput) {
				if (sampleInput % i != 0) {
					isPrime(sampleInput, ++i, lastNo);
				}
			}
			primeList.add(sampleInput);
			++sampleInput;
			isPrime(sampleInput, 2, lastNo);
		}
	}
}
