# DummKonjugator

Willkommen aus DummKonjugator! This Golang script conjugates verbs in German _present_ tense. One can conjugate multiple verbs at once by feeding the script with a comma-separated list.

## Why 'Dumm'?

It's a _dumb conjugator_ because it:
* doesn't differ regular from irregular verbs
* doesn't know if a string of letters is or isn't a verb
* treats everything ending with `-en` as a verb.

## Examples

```shell
$ go run main.go
Willkommen aus DummKonjugator.
Bitte, tippen Sie ein Verb: leben

ich        lebe
du         lebst
Sie        leben
er/sie/es  lebt
wir        leben
ihr        lebt
Sie        leben
sie        leben
```

```shell
$ go run main.go
Willkommen aus DummKonjugator.
Bitte, tippen Sie ein Verb: stiemen,glauben,fahren,sprechen,holen,haben 

ich        stieme
du         stiemst
Sie        stiemen
er/sie/es  stiemt
wir        stiemen
ihr        stiemt
Sie        stiemen
sie        stiemen

ich        glaube
du         glaubst
Sie        glauben
er/sie/es  glaubt
wir        glauben
ihr        glaubt
Sie        glauben
sie        glauben

ich        fahre
du         fahrst
Sie        fahren
er/sie/es  fahrt
wir        fahren
ihr        fahrt
Sie        fahren
sie        fahren

ich        spreche
du         sprechst
Sie        sprechen
er/sie/es  sprecht
wir        sprechen
ihr        sprecht
Sie        sprechen
sie        sprechen

ich        hole
du         holst
Sie        holen
er/sie/es  holt
wir        holen
ihr        holt
Sie        holen
sie        holen

ich        habe
du         habst
Sie        haben
er/sie/es  habt
wir        haben
ihr        habt
Sie        haben
sie        haben
```