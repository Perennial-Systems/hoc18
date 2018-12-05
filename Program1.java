public class Program1 {
	 public static void main(String []args){
         
         System.out.println(findPrimes(3,100));
  }
  
  public static long findPrimes(long a, long b){
      if(a==b){
          if(isPrime(a,2)){
             return a;    
          }
      }else if(a<=b){
          if(isPrime(a,2)){
        	  System.out.println(a);
             return findPrimes(++a,b);
          }else{
        	  findPrimes(++a,b);
          }
      }
      return findPrimes(++a,b);
  }
  public static boolean isPrime(long x, long i){
 if(x==1 && x==2){
     return true;
 }
 
 if(x%i==0 && i<x){
     return false;
 } else if(x==i){
     return true;
 }else {
     i=i+1;
     return isPrime(x,i);
 } 
 
}
}

