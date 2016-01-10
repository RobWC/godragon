# godragon
Dragon static data from Riot's League of Legends

Golang Tools to call the dragon API for League of Legends

https://developer.riotgames.com/docs/static-data

# Why?

Because Golang + Anything is better! Especially League of Legends

# Tools

## Champ Tiled Wallpaper

Create a wallpaper of any size from a champion's Icon

```
go get github.com/robwc/godragon/cmd/champtilepaper
```

Usage
```
Usage of champtilepaper:
-all
      Create wallpapers for all champs at the specified resolution
-champ string
      Specify a single or multiple champion names to create a wallpaper (-champ Teemo or -champ Teemo,Ziggs)
-height int
      Height of wallpapers (default 1080)
-output string
      Specify the output location for the wallpaper (default ".")
-width int
      Width of wallpapers (default 1920)

```

Specify the champion name
```
champtilepaper -champ Teemo
```

Specify multiple champion names
```
champtilepaper -champ Teemo,Ziggs,Rumble
```

Create your favorite team
```
champtilepaper -champ Teemo,Ahri,Rengar,Braum,Jinx
```

Specify all to create a wallpaper for every champ
```
champtilepaper -all
Creating wallpaper for Gnar at 1920x1080
Creating wallpaper for Irelia at 1920x1080
Creating wallpaper for Vi at 1920x1080
Creating wallpaper for Shyvana at 1920x1080
Creating wallpaper for Trundle at 1920x1080
Creating wallpaper for Urgot at 1920x1080
Creating wallpaper for Jinx at 1920x1080
Creating wallpaper for Katarina at 1920x1080
Creating wallpaper for Kennen at 1920x1080
Creating wallpaper for Malphite at 1920x1080
etc...
```

Specify a custom resolution
```
champtilepaper -champ Teemo -width 2880 -height 1800
Creating wallpaper for Teemo at 2880x1800
```
