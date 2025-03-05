package chpater.second;

/*
многострочный комментарий
Объявление класса Main,
который содержит код программы
*/
public class Main {

    //Определение метода main
    public static void main(String[] args) {
        boolean isActive = true;
        System.out.println(isActive);

        byte a = 3; //-128 до 127
        System.out.println(a);

        short b = 4;
        System.out.println(b); //-32768 до 32767

        int c = 4;
        System.out.println(a); //-2147483648 до 2147483647

        long d = 5;
        System.out.println(d);

        double e = 6;
        System.out.println(e);

        float f = 8.5F;
        System.out.println(f);

        char g = 'a';
        System.out.println(g);

        char ch = 100;
        System.out.println(ch);

        String text = """
                Лишь тем, кем бой за жизнь изведан,
                жизнь и свободу заслужил.
                """;

        System.out.println(text);
    }
}
