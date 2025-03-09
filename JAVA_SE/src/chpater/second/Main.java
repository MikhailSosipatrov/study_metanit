package chpater.second;

import java.math.BigDecimal;
import java.util.Scanner;

/*
многострочный комментарий
Объявление класса Main,
который содержит код программы
*/
public class Main {

    //Определение метода main
    public static void main(String[] args) {
        double d = 2.0 - 1.1;
        System.out.println(d);
        BigDecimal b = BigDecimal.valueOf(2.0 - 1.1);
        System.out.println(b);

        BigDecimal a = new BigDecimal("2.0");
        System.out.println(a);
        BigDecimal b1 = new BigDecimal("1.1");
        System.out.println(b1);
        System.out.println(a.subtract(b1));
    }
}
