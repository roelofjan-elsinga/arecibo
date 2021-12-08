# Arecibo

Arecibo: Simple text randomizer.

Transforms 

```txt
This {article|website|newspaper} will give you the {most|best|most accurate} information about this {subject|topic}.
```

into:

```txt
This article will give you the best information about this subject.
```

## Usage

```shell
./arecibo --source source-file.txt --output target-file.txt
```

Flags:
- --source, -s: Source file to generate text from
- --output, -o: Target file to write text to
- --terminal, -t: Whether to output the text to the terminal

You need to supply a source file. 
This can be any type of file.

You can choose to output the text to the terminal using the --terminal flag:

```shell
./arecibo --source source-file.txt --terminal
```