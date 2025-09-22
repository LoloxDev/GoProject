package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Contact struct {
	ID       int
	Nom      string
	Email    string
	Password string
}

var contacts = make(map[int]Contact)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nBienvenue dans le crm de gestion des contacts\n\n\n")
		fmt.Println("1. ajouter un contact\n")
		fmt.Println("2. lister les contacts\n")
		fmt.Println("3. supprimer un contact\n")
		fmt.Println("4. mettre a jour un contact\n")
		fmt.Println("5. quitter\n\n\n")

		fmt.Print("votre choix : ")
		choixStr, _ := reader.ReadString('\n')
		choix, err := strconv.Atoi(strings.TrimSpace(choixStr))
		if err != nil {
			fmt.Println("erreur")
			continue
		}

		switch choix {
		case 1:
			ajouterUnUtilisateur(reader)
		case 2:
			listerLesUtilisateurs()
		case 3:
			supprimerUnUtisateur(reader)
		case 4:
			mettreAJourUnUtilisateur(reader)
		case 5:
			fmt.Println("au revoir")
			return
		default:
			fmt.Println("saissez 1-5 svp")
		}
	}
}

func ajouterUnUtilisateur(reader *bufio.Reader) {
	fmt.Print("id : ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("id incorrect")
		return
	}
	if _, ok := contacts[id]; ok {
		fmt.Println("un contact a déja choisis cet id")
		return
	}

	fmt.Print("votre nom : ")
	nom, _ := reader.ReadString('\n')

	fmt.Print("votre email : ")
	email, _ := reader.ReadString('\n')

	fmt.Print("votre mot de passe : ")
	password, _ := reader.ReadString('\n')

	contacts[id] = Contact{id, strings.TrimSpace(nom), strings.TrimSpace(email), strings.TrimSpace(password)}
	fmt.Println("contact ajouté !")
}

func listerLesUtilisateurs() {
	if len(contacts) == 0 {
		fmt.Println("aucun contact n'a encore ete enregistré")
		return
	}
	for _, c := range contacts {
		fmt.Printf("%d - %s (%s)\n", c.ID, c.Nom, c.Email)
	}
}

func supprimerUnUtisateur(reader *bufio.Reader) {
	fmt.Print("id a supprimer: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("id pas bon")
		return
	}
	if _, ok := contacts[id]; ok {
		delete(contacts, id)
		fmt.Println("contact suprimé")
	} else {
		fmt.Println("contact introuvable..")
	}
}

func mettreAJourUnUtilisateur(reader *bufio.Reader) {
	fmt.Print("id a modifier: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("id pas correct")
		return
	}
	c, ok := contacts[id]
	if !ok {
		fmt.Println("contact introuvable")
		return
	}

	fmt.Print("nouveau nom ? ")
	nom, _ := reader.ReadString('\n')
	fmt.Print("nouvel email ? ")
	email, _ := reader.ReadString('\n')

	c.Nom = strings.TrimSpace(nom)
	c.Email = strings.TrimSpace(email)
	contacts[id] = c
	fmt.Println("le contact a bien été mis a jour")
}
