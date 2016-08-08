# 5주차 과제

## 과제 설명
5주차 과제인 'Simple Server'는 GoLang과 Go루틴을 이용해 제작하는 간단한 채팅 서버이다.  
이 과제를 통해 Go언어의 꽃인 Go루틴을 이용한 실질적인 서버개발을 경험하게 된다.

## 진행 단계

### 프로그램 설정
- 자원의 제약이 없다고 가정했을 시, 서버는 지정한 포트로 접속하는 클라이언트들을 무제한으로 받아들일 수 있어야 한다.
- 병렬적인 동작이 필요한 모든 경우에는 Go루틴을 사용하여 처리한다.
- select/case문을 한 번 이상 사용하여 구현한다.
- 클라이언트에게 받은 데이터는 해당 클라이언트를 포함한 모든 클라이언트에게 송신되어야 한다.

### 프로그램 구조
- Client와 ChatRoom 이라는 struct가 아래와 같이 주어진다.

```
type Client struct {
	input chan string
	output chan string
	receiver *bufio.Reader
	sender *bufio.Writer
}

type ChatRoom struct {
	clientList []*Client
	newClient chan net.Conn
	input chan string
}
```

- 필요하면 위 struct를 어느정도 수정하여 사용하여도 괜찮다.
- 모든 struct의 객체는 New[struct명] 메서드를 이용하여 받아와야 한다. (ex. NewClient(), NewChatRoom())
- Client와 ChatRoom 객체는 모두 Listen() 메서드를 가지고 있으며, New[struct명] 메서드 내에서 Go루틴을 통해 객체가 소멸될 때까지 실행되어야 한다. 이 메서드 내에서 Client의 경우 입/출력을 진행하고, ChatRoom의 경우 Join과 메세지 전달을 수행한다.
- input/output을 위한 메서드를 따로 구현하는 것은 개인의 자유로 한다.
- 클라이언트를 식별할 수 있는 이름은 굳이 필요하지 않으나, 원하면 넣어도 상관없다.


## 실행 결과

### 서버
```
skyclad@Skyclad-Server:~/src/goserver/server$ ./server
[*] Server started
[*] Make new client
[*] Make new client
Hello, I'm Client1
Hello, Nice to meet you
1111
2222
```


### 클라이언트 1
```
skyclad@Skyclad-Server:~$ nc localhost 25252
Hello, I'm Client1
Hello, I'm Client1
Hello, Nice to meet you
1111
1111
2222
```

### 클라이언트 2
```
skyclad@Skyclad-Server:~$ nc localhost 25252
Hello, I'm Client1
Hello, Nice to meet you
Hello, Nice to meet you
1111
2222
2222
```

(자신이 전송한 메시지도 서버에 의해 자신에게 재전송되므로 두 번 보여진다)