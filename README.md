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

## Installation

### API key

Before starting to use the tool, you need to get yourself an OpenAPI API key - that can be done [here](https://openai.com/product).

With your API key at hand, you need to set up the config file, in which to store it:

```bash
cd Prompt
 echo "api_key: <your key>" > config.yaml 
```

Note that by adding a whitespace before the echo command, your command will not appear in your shell history - this is recommended, given the use of an API key.
If you prefer, you can also use any editor/IDE to create the `config.yaml`file and input your key into it. The only relevant thing is that it looks as follows:

```yaml
api_key: <your key>
```

### Building the binary

Run the following to clone the repository, build and move the binary:

```bash
git clone https://github.com/patrick-men/GoGPT.git
cd GoGPT
go build 
cp GoGPT /usr/local/bin
```

With this done, you simply need to restart your shell by running `zsh`, `bash`, `sh` or whichever shell you're using.

Now that you have the binary, you can use the binary with flags, or interactively:

#### Flags

To use the binary with flags, simply add your prompt after the binary:

```bash
GoGPT "your prompt here"
```

This is recommended in cases, where you want to implement the tool into any automation, such as other tools.

#### Interactive

To run the tool interctively, you simply need to call it with `GoGPT`, and then write your prompt. 

## Use different Model

The default model is GPT3.5 Turbo. If you want to use another model, you can find the possible models [here](https://github.com/sashabaranov/go-openai/blob/a09cb0c528c110a6955a9ee9a5d021a57ed44b90/completion.go#L20).

If you want to change it, you simply need to change the values to the corresponding model in this line here:
![](https://github.com/patrick-men/GoGPT/blob/5b1ee4e571bd1dbea246905bde71caaa94434403/Prompt/api.go#L40)