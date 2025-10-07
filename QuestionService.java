import java.util.Scanner;

public class QuestionService {
    Question[] questions = new Question[3];

    QuestionService() {
        questions[0] = new Question(0, 
        "What is the best PPL?",
        new String[] {"Java", "Python", "CPP"}, 
        "Python");

        questions[1] = new Question(1, 
        "Who is handsome?",
        new String[] {"Chan", "BC", "SYB"}, 
        "Chan");

        questions[2] = new Question(2, 
        "The most beautiful girl in the world?",
        new String[] {"Taylor", "Sydney Sweeney", "Elllen Fanning"}, 
        "Taylor");

    }

    public void startQuiz() {
        System.out.println("Welcome to my Quiz");

        int score = 0;
        String userAnswer;

        for (Question q : this.questions) {
            System.out.println("Question no: " + q.getId());
            System.out.println(q.getQuestion());
            System.out.println("Choose one:");

            for (String option : q.getOptions()) {
                System.out.println("- " + option);
            }

            Scanner sc = new Scanner(System.in);
            userAnswer = sc.nextLine();
            
            if (userAnswer.equals(q.getAnswer())) {
                System.out.println("Your answer:" + userAnswer + " is correct!");
                score += 1;
            } else {
                System.out.println("Your answer:" + userAnswer + " is not correct!");
                score -= 1;
            }

            userAnswer = "";
            sc.close();
        }

        if (score < 0) {
            score = 0;
        }

        System.out.println("Your score is: " + score);
        System.out.println("Good job");
        
    }
}