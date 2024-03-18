package pflag

import (
	"bytes"
	"io"
	"testing"
)

const expectedOutput = `      --long-form    Some description
      --long-form2   Some description
                       with multiline
  -s, --long-name    Some description
  -t, --long-name2   Some description with
                       multiline
`

func setUpPFlagSet(buf io.Writer) *FlagSet {
	f := NewFlagSet("test", ExitOnError)
	f.Bool("long-form", false, "Some description")
	f.Bool("long-form2", false, "Some description\n  with multiline")
	f.BoolP("long-name", "s", false, "Some description")
	f.BoolP("long-name2", "t", false, "Some description with\n  multiline")
	f.SetOutput(buf)
	return f
}

func TestPrintUsage(t *testing.T) {
	buf := bytes.Buffer{}
	f := setUpPFlagSet(&buf)
	f.PrintDefaults()
	res := buf.String()
	if res != expectedOutput {
		t.Errorf("Expected \n%s \nActual \n%s", expectedOutput, res)
	}
}

func setUpPFlagSet2(buf io.Writer) *FlagSet {
	f := NewFlagSet("test", ExitOnError)
	f.Bool("long-form", false, "Some description")
	f.Bool("long-form2", false, "Some description\n  with multiline")
	f.BoolP("long-name", "s", false, "Some description")
	f.BoolP("long-name2", "t", false, "Some description with\n  multiline")
	f.StringP("some-very-long-arg", "l", "test", "Some very long description having break the limit")
	f.StringP("other-very-long-arg", "o", "long-default-value", "Some very long description having break the limit")
	f.String("some-very-long-arg2", "very long default value", "Some very long description\nwith line break\nmultiple")
	f.SetOutput(buf)
	return f
}

const expectedOutput2 = `      --long-form                      Some description
      --long-form2                     Some description
                                         with multiline
  -s, --long-name                      Some description
  -t, --long-name2                     Some description with
                                         multiline
  -o, --other-very-long-arg <string>   Some very long description having
                                       break the limit (default
                                       "long-default-value")
  -l, --some-very-long-arg <string>    Some very long description having
                                       break the limit (default "test")
      --some-very-long-arg2 <string>   Some very long description
                                       with line break
                                       multiple (default "very long
                                       default value")
`

func TestPrintUsage_2(t *testing.T) {
	buf := bytes.Buffer{}
	f := setUpPFlagSet2(&buf)
	res := f.FlagUsagesWrapped(80)
	if res != expectedOutput2 {
		t.Errorf("Expected \n%q \nActual \n%q", expectedOutput2, res)
	}
}

func setUpPFlagSet3(buf io.Writer) *FlagSet {
	f := NewFlagSet("test", ExitOnError)
      f.StringP("birth-date", "b", "2000-01-01", "the `birth date` of our main character") // results in <birth_date>
      f.String("event", "birthday party", "choose from `birthday party, wedding or graduation`") // results in {birthday party|wedding|graduation}
	f.StringArray("friend", []string{"Alice", "Bob"}, "the friends of our main actor, in the format of `name1,...`") // results in <name1,...>
	f.StringP("name", "l", "Alice", "the `name` of our main actor") // results in <name>
	f.String("place", "park", "choose from `park, home, office`")  // results in {park|home|office}
	f.SetOutput(buf)
	return f
}

const expectedOutput3 = `  -b, --birth-date <birth_date>                     the birth date of our
                                                    main character
                                                    (default "2000-01-01")
      --event {birthday party|wedding|graduation}   choose from birthday
                                                    party, wedding or
                                                    graduation (default
                                                    "birthday party")
      --friend <name1,...>                          the friends of our
                                                    main actor, in the
                                                    format of name1,...
                                                    (default [Alice,Bob])
  -l, --name <name>                                 the name of our main
                                                    actor (default "Alice")
      --place {park|home|office}                    choose from park,
                                                    home, office (default
                                                    "park")
`

func TestPrintUsage_3(t *testing.T) {
	buf := bytes.Buffer{}
	f := setUpPFlagSet3(&buf)
	res := f.FlagUsagesWrapped(80)
	if res != expectedOutput3 {
		t.Errorf("Expected \n%q \nActual \n%q", expectedOutput3, res)
	}
}
