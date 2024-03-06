# API Quiz App

API ini sudah di hosting VM dan juga menggunakan Domain http://api-quiz-arras.my.id:8080
Berikut Adalah Dokumentasi Penggunaan API ini:
- POST Login http://api-quiz-arras.my.id:8080/api/user/login 
{
    "email": "arras@email.com",
    "password": "arras"
}

- POST Register http://api-quiz-arras.my.id:8080/api/user/register
{
    "username": "arras",
    "email": "arras@email.com",
    "password": "arras"

}

- POST Created Quiz with Auth Bearer Token http://api-quiz-arras.my.id:8080/api/quiz/created-quiz
{
    "title": "ini title quiz",
    "description": "ini description pembuatan dari postman ini description pembuatan dari postmann",
    "startedAt": "2024-02-22T21:41:49Z",
    "finishedAt": "2024-02-29T21:41:49Z"
}

- PUT Update Quiz with Auth Bearer Token http://api-quiz-arras.my.id:8080/api/quiz/updated-quiz/:id
{
    "title": "Create Quiz Update",
    "description": "Test Create Quiz",
    "startedAt": "2024-02-22T21:41:49Z",
    "finishedAt": "2024-02-27T21:41:49Z"
}

- GET All Quiz with Auth Bearer Token http://api-quiz-arras.my.id:8080/api/quiz/
- GET All Quiz Active with Auth Bearer Token http://api-quiz-arras.my.id:8080/api/quiz/is-active
- GET Quiz by ID with Auth Bearer Token http://api-quiz-arras.my.id:8080/api/quiz/:id
- GET Question by ID Quiz with Auth Bearer Token http://api-quiz-arras.my.id:8080/api/question/questions/:quiz_id
- POST Create Question to Quiz with Auth Bearer Token http://api-quiz-arras.my.id:8080/api/question/created-question/:id
[
    {
        "question": "Pertanyaan 1",
        "true_answer": "B",
        "options": {
            "A": "Pilihan A",
            "B": "Pilihan B",
            "C": "Pilihan C",
            "D": "Pilihan D"
        }
    },
    {
        "question": "Pertanyaan 2",
        "true_answer": "B",
        "options": {
            "A": "Pilihan A",
            "B": "Pilihan B",
            "C": "Pilihan C",
            "D": "Pilihan D"
        }
    }

]

- POST Student Submit Answer with Auth Bearer Token http://api-quiz-arras.my.id:8080/api/submit-answer/quizzes/:quizID/submit-answers
[
    {
        "questionID": 1,
        "answer": "A"
    },
    {
        "questionID": 2,
        "answer": "A"
    }
  
]

- GET Score User by Quiz with Auth Bearer Token http://api-quiz-arras.my.id:8080/api/submit-answer/api/user/quiz/:quizID/total-score
- DELETE Quiz with Auth Bearer Token http://api-quiz-arras.my.id:8080/api/quiz/:id


