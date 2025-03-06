package chpater.second;

import java.util.Scanner;

/*
многострочный комментарий
Объявление класса Main,
который содержит код программы
*/
public class Main {

    //Определение метода main
    public static void main(String[] args) {
        System.out.println("Hello!");
        System.out.print("Bye!");
        System.out.printf("x=%s", "stroka");
        Scanner in = new Scanner(System.in);
        System.out.println("Input a number: ");
        int x = in.nextInt();
        System.out.println("Your number is: " + x);
        in.close();
    }
}
