# Mini-CRM CLI

Un gestionnaire de contacts simple et efficace en ligne de commande, écrit en Go.
Ce projet illustre les bonnes pratiques de développement Go :

- Architecture en packages découplés
- Injection de dépendances via interfaces
- CLI professionnelle avec Cobra
- Configuration externe avec Viper
- Plusieurs couches de persistance (GORM/SQLite, JSON, Memory)

## Fonctionnalités

- Ajouter un contact (id, nom, email, mot de passe)
- Lister tous les contacts
- Modifier un contact existant
- Supprimer un contact
- Gestion multi-backends (SQLite, JSON, mémoire)

## Installation

```bash
git clone https://github.com/LoloxDev/GoProject/tree/TpJeudi
cd GoProject
go mod tidy
```

## Configuration

L’application lit le fichier config.yaml pour choisir le type de stockage.

Exemple :

```bash
storage:
  type: gorm   # valeurs possibles : gorm | json | memory

  gorm:
    path: data/crm.db

  json:
    path: data/contacts.json

  memory: {}
```

- gorm → stockage persistant SQLite dans un fichier .db
- json → stockage persistant JSON lisible par l’humain
- memory → stockage éphémère (perdu à chaque exécution)

## Utilisation

Toutes les commandes se lancent via :

```bash
go run . <commande> [options]
```

Ajouter un contact

```bash
go run . add --name "Jane Doe" --email jane@example.com --password secret
```

Lister les contacts

```bash
go run . list
```

Mettre à jour un contact

```bash
go run . update --id 1 --name "Jane D." --email jane.d@example.com
```

Supprimer un contact

```bash
go run . delete --id 1
```


<img width="1046" height="341" alt="logo_efrei_web_bleu" src="https://github.com/user-attachments/assets/f8838509-0eaa-4f0d-89dc-f55e59cb8c79" />
