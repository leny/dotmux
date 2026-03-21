# dotmux

> Dotfiles &amp; scripts for tmux

* * *

## Install

    mkdir -p ~/.config/tmux-plugins
    git clone https://github.com/tmux-plugins/tmux-resurrect ~/.config/tmux-plugins/tmux-resurrect
    git clone https://github.com/leny/dotmux.git ~/.dotmux
    ln -sfv ~/.dotmux/tmux.conf ~/.tmux.conf

## Alerts

Système d'alertes cross-session pour tmux. Affiche un indicateur dans la status bar quand une session nécessite de l'attention, avec un menu interactif pour naviguer vers les sessions en attente.

### Utilisation

```bash
# Déclencher une alerte pour la session courante
tmux-alert set

# Déclencher une alerte pour une session spécifique
tmux-alert set my-session

# Effacer l'alerte de la session courante
tmux-alert clear

# Effacer l'alerte d'une session spécifique
tmux-alert clear my-session

# Effacer toutes les alertes
tmux-alert clear-all
```

### Intégration tmux

- **Status bar** : les sessions en alerte apparaissent en noir sur fond jaune (ex: `⚡ my-session (3m)`)
- **`prefix + v`** : ouvre un menu pour naviguer vers une session en alerte
- **`prefix + space` → Alerts** : même menu via le menu Commands
- **Auto-clear** : l'alerte est automatiquement effacée quand on switch vers la session concernée

### Intégration dans des scripts/hooks

Le système est générique et peut s'intégrer n'importe où. Quelques exemples :

```bash
# Dans un hook Claude Code (settings.json)
# → set quand Claude attend de l'input, clear quand il reprend
tmux-alert set "$SESSION_NAME"
tmux-alert clear "$SESSION_NAME"

# Dans un script de build long
./build.sh && tmux-alert set

# Dans un watcher qui surveille un processus
my-long-task; tmux-alert set

# Depuis un cron ou un process distant (via SSH)
ssh myhost 'tmux-alert set build-server'
```

### Configuration

| Variable | Défaut | Description |
|---|---|---|
| `TMUX_ALERT_DIR` | `/tmp/tmux-alerts` | Répertoire des marqueurs d'alerte |

Les marqueurs sont de simples fichiers contenant un timestamp epoch, nommés d'après la session tmux (les `/` sont encodés en `%2F`). Cela permet de les inspecter/manipuler facilement :

```bash
# Voir les alertes actives
ls /tmp/tmux-alerts/

# Inspecter un marqueur
cat /tmp/tmux-alerts/my-session        # → 1709398400
cat /tmp/tmux-alerts/work%2Fmy-project # → session "work/my-project"
```

* * *

April 2018
