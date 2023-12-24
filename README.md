# GoGPT

This is a CLI-Tool written in Golang that talks to the OpenAI API. The goal is to achieve a way to communicate with the GPT-Models, whilst staying in the console. Furthermore, this allows communication with the more advanced models, though at a small cost:

| Model                  | Price input / 1k tokens | Price output / 1k tokens |
| ---------------------- | ----------------------- | ------------------------ |
| gpt-3.5-turbo-1106     | $0.00010                | $0.00020                 |
| gpt-3.5-turbo-instruct | $0.00015                | $0.00020                 |

This model is already more advanced than the freely available GPT 3, tho there is an even more potent one available. This does, however, come at a noticable increase in cost:

| Model     | Price input / 1k tokens | Price output / 1k tokens |
| --------- | ----------------------- | ------------------------ |
| gpt-4     | $0.03                   | $0.06                    |
| gpt-4-32k | $0.06                   | $0.12                    |

Given that the GPT 3.5 model is already a notable increase, this is the chosen way to go from here on out.

> This needs to be checked, but in case the way to communicate with the API of a different model doesn't differ too much, this tool can be forked and adapted.
> Best case, only the API key can be changed appropriately.

## Usage

The main goal of this tool is to be used as a CLI tool, meaning you don't have to leave the terminal/IDE to use ChatGPT. It can, and will, however, be integrated into further workflows:

- Use it as a backend for your web application
- Extend the interactive part with your own code/tool
- Integrate it into existing tools
- etc.

## Installation

### Clone the repository

```bash
git clone https://github.com/patrick-men/GoGPT.git
```

### Binary

Install the binary from the release, and then move it to the binary directory:

```bash
mv <> /usr/local/bin
```
