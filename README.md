
# Number Guessing Game

Sample solution for the [number-guessing-game](https://roadmap.sh/projects/number-guessing-game) challenge from [roadmap.sh](https://roadmap.sh).

## Introduction

Welcome to the Number Guessing Game! This is a simple command-line game written in Go where you guess a randomly generated number between 1 and 100. The game features three difficulty levels: Easy, Medium, and Hard, each with a different number of attempts to guess the correct number.

## How to Run

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/your-username/number-guessing-game-go.git
   cd number-guessing-game-go
   ```

2. **Build and Run the Game**:

   Ensure you have Go installed on your system. Then, execute the following:

   ```bash
   go run main.go
   ```

## How to Play

1. Run the game using the instructions in the "How to Run" section.
2. You will see a welcome message and a prompt to select the difficulty level:
   - Enter `1` for Easy.
   - Enter `2` for Medium.
   - Enter `3` for Hard.
3. After selecting the difficulty level, you will be prompted to guess a number between 1 and 100.
4. Enter your guesses one at a time.
   - If your guess is incorrect, the game will provide a hint: whether the correct number is higher or lower than your guess.
   - If your guess is correct, the game will congratulate you!
5. If you run out of attempts, the game will reveal the correct number.

## Difficulty Levels

- **Easy**: 10 attempts to guess the correct number.
- **Medium**: 5 attempts to guess the correct number.
- **Hard**: 3 attempts to guess the correct number.

## Technical Details

This game is implemented in Go. The random number generation and user input handling are simple and efficient, making it a lightweight application.

### Requirements

- Go (version 1.16 or higher)

### Code Structure

- `main.go`: The main entry point of the application, containing all the game logic.

## Contributing

If you'd like to contribute to this project, please fork the repository and submit a pull request with your changes.

### Contributing Guidelines

- Follow Go best practices for clean and idiomatic code.
- Ensure your code is well-documented and tested before submitting a pull request.
- Keep your changes concise and focused on a single issue or feature.

## Authors

- [Mehul](https://github.com/mehulnaik16)

## Acknowledgments

- Thanks to the Go community for their resources and tools.
- Thanks to Roadmap.sh for inspiring this project.
- Thanks to all contributors who make open-source projects possible!
