# 3주차 과제

## 과제 설명
3주차 과제인 'Custom Creature'는 interface와 struct를 이용하여 종족을 만들고, 메서드를 구현하도록 하는 프로그램이다.  
이 과제를 통해 Go언어의 struct, interface에 대해 이해하고 다룰 수 있게 된다.  

## 진행 단계

### 프로그램 설정
- "fmt" 외의 다른 모듈은 import하지 않는다(단, 수학적 계산 등 프로그래밍에 큰 영향을 주지 않는 경우는 허가한다).  
- 종족명, 객체의 name, atk, spcatck 등 멤버 변수는 임의로 설정한다.

### interface 구현하기
- 아래와 같은 interface가 주어진다.
```
type Creature interface {
	SelfIntroduction()
	NormalAttack()
	SpecialAttack()
}
```
- SelfIntoruction()에서는 string 타입의 name 변수를 이용해 자기소개를 한다.
	- 자기소개는 종족마다 다르며, 형식도 달라져야 한다. 자세한 것은 아래의 예시 참조.
- NormalAttack()에서는 int 타입의 atk 변수를 이용해 무기명과 Damage를 출력한다.
- SpecialAttack()에서는 int 타입의 spcatk 변수를 이용해 기술 또는 무기명과 Damage를 출력한다.

### struct 구현하기
- 기본적인 struct의 형식은 아래를 따른다.
```
type 종족 struct {
	name string
	atk int
	spcAtk int
}
```
- 각 struct마다 **포인터 리시브를 이용한** 메서드를 구현하여야 하며, 자세한 내용은 **interface 구현하기** 를 참고한다.

### 출력하기
- 출력을 위한 메서드를 따로 제작한다.
	- 인자로 Creature 타입의 interface를 받고, 내부에서 interface의 메서드를 실행시켜 출력한다.
	- 메서드명은 CreatureInfo이며, 리턴하지 않는다.


## 실행 결과
```
<< Creature Information Service >>
I'm a Human, My name is King Arthur
Normal Attack: Sword
Damage 100
Special Attack: Excalibur
Damage 1000
===================================
Nice to meet you, My name is Serenial and I'm a Elf.
Normal Attack: Bow
Damage 50
Special Attack: FireArrow
Damage 300
===================================
Hello! My name is Excelhand , a Dwarf
Normal Attack: Hammer
Damage 150
Special Attack: Volcanic Eruption
Damage 750
```