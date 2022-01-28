class: center, middle

# Bowling Kata

---

# How to score Bowling

http://en.wikipedia.org/wiki/Ten-pin_bowling

![Bowling Frame](http://upload.wikimedia.org/wikipedia/commons/thumb/6/61/Bowlstrike.PNG/330px-Bowlstrike.PNG)

---


# External API

- call roll for each roll
- calculate score at the end
- Assume good inputs (rolls always equal a full game)
- Start with a command line app, evolve to a class in a library


---


#Gutter Game

If all 20 rolls are 0 (gutter ball), the score should be 0

```bash
> ./bowl 0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0
0
```


---

#Simple games

If all 20 rolls are 1 the score should be 20

```
> ./bowl 1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1
1
```

---

# Sum games

If all rolls are less than or equal to 4, the score is the sum of the rolls

- roll '3' 10 times
- roll '4' 10 times
    - score should be 70

---

#Refactor Tests

## Always be stable!

Use *Working Code* as tests for test code

- Perhaps change manually typed array or rolls into a generator with a loop
- Does this belong in Production code OR Test code?
- Use the production code as the test code when refactoring tests

---

# Game with a single spare

if we roll a 5 three times, and then all zeros the score should be 20

- roll '5' 3 times
- roll '0' 17 times
    - Score should be 20

You may need to remove this test, refactor, then procede

---

#Possible object model
.omodel[![object model](https://raw.github.com/pixelhandler/vagrant-dev-env/bowling/www/app/images/frame_class_next.png)]
---

# Game with a single strike

if we roll a Strike , then a 3 two times, and then all zeros the score should be 16

- roll '10' one time
- roll '3' two times
- roll '0' 16 times
    - Score should be 22

---

# Game with mixed strikes and spares

Rolls:
3,4,6,4,7,1,4,5,10,4,6,10,10,10,5,5,3 = 169

---

# New Requirement!

## Bad input!
- Throw an exception if score is called and you don't have a complete game

---

# New Requirement!

## Per frame scoring

- be able to call score at any time and return the current score
- return known score if current frame cannot be scored


---

# New Requirement!

## Kids bowling

- GutterBalls are not allowed
- All rolls of 0 are re-rolled
- Kid bowling will be indicated by the caller (constructor arg?)

 
---

# Klingon Bowling

- Each frame gets 3 rolls
- the lowest roll is considered ghojwI' (“student”)
- two consecutive 0 rolls earns bIj (“punishment”) of -10

---

# Temporal Bowling

- Final Score is multiplied by the minutes past the hour based on when the function was called

---

# Credits

- Images and inspiration from this post:
  http://pixelhandler.com/posts/bowling-game-kata-using-mocha-bdd-test-framework-and-yeoman

- Unclebob's original post
  http://butunclebob.com/ArticleS.UncleBob.TheBowlingGameKata