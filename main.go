package main

/* Le programme a un système de langueB, de niveau, de nombre de partis souhaités et affiche les lettres déjà utilisées.
Pour la version française, on s'est servi des fichiers texte,
et pour la version anglaise on se sert de l'API d'un site qui crew des listes de mots de la longueur demandé.
*/

var (
	lettreU string
	score   string
	coupF   = 0
)

func main() {
	pendu()
}
