# Go-reloaded

## Fichiers et Fonctions

1. **`main.go`**

    Ce fichier contient le point d'entrée principal de l'application et orchestre le processus d'édition de texte.

    - **`main()`**: La fonction principale lit les arguments de la ligne de commande, lit le contenu du fichier d'entrée, applique les opérations d'édition de texte et écrit le texte édité dans le fichier de sortie.

2. **`HexBin.go`**

    Contient des fonctions pour convertir les représentations hexadécimales et binaires en décimal.

    - **`EditHexBin(text string) string`**: Remplace les représentations hexadécimales et binaires dans le texte d'entrée par leurs équivalents décimaux.

3. **`Case.go`**

    Implémente diverses fonctionnalités de modification de cas.

    - **`EditCase(text string) string`**: Effectue des transformations en majuscules, minuscules et capitalisées sur le texte d'entrée en fonction de motifs spécifiés.
    - **`processMultipleCase(text string, Trim *regexp.Regexp, caseFunc func(string) string) string`**: Gère les modifications de cas pour les motifs avec un nombre spécifié de mots et de cas.

4. **`Punctuation.go`**

    Manipule la ponctuation dans le texte d'entrée pour assurer un espacement et un placement corrects.

    - **`EditPunctuation(text string) string`**: Gère les signes de ponctuation dans le texte d'entrée, assurant un espacement et un placement corrects.

5. **`Quote.go`**

    Nettoie les guillemets dans le texte d'entrée pour assurer un espacement correct.

    - **`EditQuote(text string) string`**: Nettoie les guillemets et assure un espacement correct autour d'eux.

6. **`AN.go`**

    Implémente la fonctionnalité de conversion de "A" en "An".

    - **`EditAN(text string) string`**: Remplace les occurrences de "a" par "an" avant les mots commençant par une voyelle ou 'h'.


