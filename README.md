# Go-reloaded

## Fichiers et Fonctions

1. **`main.go`**

    - **`main()`**: Cette fonction est la fonction principale. Elle vérifie les arguments de la ligne de commande, lit le contenu du fichier d'entrée, applique les opérations d'édition de texte et écrit le texte édité dans le fichier de sortie.

2. **`HexBin.go`**

    - **`EditHexBin(text string) string`**: Cette fonction remplace les représentations hexadécimales et binaires dans le texte d'entrée par leurs équivalents décimaux.

3. **`Case.go`**

    - **`EditCase(text string) string`**: Cette fonction effectue des transformations en majuscules, minuscules et capitalisées sur le texte d'entrée en fonction de motifs spécifiés.
    - **`processMultipleCase(text string, Trim *regexp.Regexp, caseFunc func(string) string) string`**: Cette fonction gère les modifications de cas pour les motifs avec un nombre spécifié de mots à modifier.

4. **`Punctuation.go`**

    - **`EditPunctuation(text string) string`**: Cette fonction gère les signes de ponctuation dans le texte d'entrée, assurant un espacement et un placement corrects.

5. **`Quote.go`**

    - **`EditQuote(text string) string`**: Cette fonction nettoie les guillemets et assure un espacement correct autour d'elles.

6. **`AN.go`**

    - **`EditAN(text string) string`**: Cette fonction remplace les occurrences de "a" par "an" avant les mots commençant par une voyelle ou 'h'.


