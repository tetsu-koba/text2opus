# text2opus
Convert text to opus audio file with Google text-to-speech API

## Build

```
$ go build
$ (cd list_voices && go build)
```

## Usage

Make sure the $GOOGLE_APPLICATION_CREDENTIALS environment variable is set correctly.
```
$ cat $GOOGLE_APPLICATION_CREDENTIALS
{
    some json contents ...
}    
```

```
$ ./text2opus 
2023/02/25 16:20:17 Usage:./text2opus input_text_file output_opus_file [voice_name]
```

example


```
$ ./text2opus hello.txt hello.opus
```

You can specify a voice name as the 3rd argument.
```
$ ./text2opus hello.txt hello_sa.opus en-US-Standard-A
```

You can see the list of voices with the `list_voices' command.

```
$ ./list_voices/list_voices 
2023/02/25 16:31:37 Usage:./list_voices/list_voices languageCode (ex. 'en-US', '' means all)
```

```
$ ./list_voices/list_voices en-US |head
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

If you want to see the full list of voices, specify '' (zero length string) as the language code.

```
$ ./list_voices/list_voices '' |wc -l
407
```
