# text2opus
Convert text to opus audio file with Google text-to-speech API

## Install

```
$ go install github.com/tetsu-koba/text2opus@latest
```

It is expected to be installed in $GOPATH/bin or $HOME/go/bin.  It has been tested on Linux x86_64 and Mac AArch64 platforms.

## Usage

Ensure the `$GOOGLE_APPLICATION_CREDENTIALS` environment variable is set correctly.

```
$ cat $GOOGLE_APPLICATION_CREDENTIALS
{
    some json contents ...
}    
```

```
$ text2opus
2023/02/25 20:44:36 Usage:
2023/02/25 20:44:36   text2opus input_text_file output_opus_file [voice_name]
2023/02/25 20:44:36   text2opus -l [languageCode (ex. 'en-US')]
```

To convert a text file into opus audio file, simpley run:

```
$ text2opus hello.txt hello.opus
```

You can specify a voice name as the 3rd paramater.

```
$ text2opus hello.txt hello_sa.opus en-US-Standard-A
```

To see the list of available voices, use the -l option.

```
$ text2opus -l en-US |head
language_codes:"en-US"  name:"en-US-Standard-A"  ssml_gender:MALE  natural_sample_rate_hertz:24000
language_codes:"en-US"  name:"en-US-Standard-B"  ssml_gender:MALE  natural_sample_rate_hertz:24000
language_codes:"en-US"  name:"en-US-Standard-C"  ssml_gender:FEMALE  natural_sample_rate_hertz:24000
language_codes:"en-US"  name:"en-US-Standard-D"  ssml_gender:MALE  natural_sample_rate_hertz:24000
language_codes:"en-US"  name:"en-US-Standard-E"  ssml_gender:FEMALE  natural_sample_rate_hertz:24000
language_codes:"en-US"  name:"en-US-Standard-F"  ssml_gender:FEMALE  natural_sample_rate_hertz:24000
language_codes:"en-US"  name:"en-US-Standard-G"  ssml_gender:FEMALE  natural_sample_rate_hertz:24000
language_codes:"en-US"  name:"en-US-Standard-H"  ssml_gender:FEMALE  natural_sample_rate_hertz:24000
language_codes:"en-US"  name:"en-US-Standard-I"  ssml_gender:MALE  natural_sample_rate_hertz:24000
language_codes:"en-US"  name:"en-US-Standard-J"  ssml_gender:MALE  natural_sample_rate_hertz:24000
```

To view the full list of voices, just specify -l.

```
$ text2opus -l |wc -l
407
```

## Update

To update text2opus, remove the executable file and then run go install again.

```
$ rm $(which text2opus)
$ go install github.com/tetsu-koba/text2opus@latest
```
