import java.util.*;

public class Program1{

	public static ArrayList<Integer> primeNumbers = new ArrayList<Integer>();

	public static void main(String[] args){
		findPrime(1, 25);
		System.out.println("Prime Numbers" + primeNumbers.toString());
	}


	public static void findPrime(int start, int end){
		if(start > end){
			return;
		} else {
			if(isPrime(start, start -1)){
				primeNumbers.add(start);
			} 
			findPrime(start+1, end);
		}
	}
	
	public static boolean isPrime(int n, int d){
		if(n == 1) return true;
		return d > 1 ? (n%d == 0 ? false : isPrime(n, d-1)) : true;
	}
}