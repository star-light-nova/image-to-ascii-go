<!-- Note: This is ok, because HTML thinks that `<` and `>` are tags. -->
<!-- So, therefore used its own notation. (&gt; &lt;) -->
<pre>
     "v)7(((JJz    rF(7J)))((()))J7(Fr        `)||^          `?JJ(((7)T'  `?J7(((7)T.     
   . jgBUOOOAg1_ ` 8$9G8$B#OOG#B$8Gd$D      . 5h4$3 .        =egAOOpUgBn  +agAOOpAg8) `   
   ` Lw!bgggy-).  :ygggl+-zggB*'!ugggw^     - |ggg5_         -|_Egggd=]z  "}^qggg6,x;     
        gggK{   . tHgY].  fmgOF  .q2g8u `    'l#dggm .          3Xggm        jRg$V        
      : ]$g6_ ' - :gHE  . n8gG1 .  p$gB -  . hg-oggh .        ' ^pgRY "    ' TAgAu "      
      . q$gh&gt;   - tgR ' . [8gG3 .`. +gJ -  - 5g^F#gj;         ` +OgDa .    ` (AgKo .      
      . w$gh;   ` &gt;gV . ` [8gGI ` - 'gR .   ,jUs {gg' -         &lt;OgRy .    ` JAgKe .      
      . w$gh;     &lt;q.   ` [8gGI `   ^p.   - 8g,  ]gg# `         &lt;OgRy .    ` JAgKe .      
      . w$gh;           ` [8gGI `         . Ogr  emgS! `        &lt;OgRy .    ` JAgKe .      
      . w$gh;     -.    ` [8gGI `    ..    ;kVr ` Lggr -        &lt;OgRy .    ` JAgKe .      
      . w$gh;           ` [8gGI `        - =gR `- egg&lt;          &lt;OgRy .    ` JAgKe .      
      . w$gh;           ` [8gGI `        ` "g_ '- 5#g4| `       &lt;OgRy .    ` JAgKe .      
      . w$gh;           ` [8gGI `       ` TO9+     /ggs -       &lt;OgRy .    ` JAgKe .      
      . w$gh;           ` [8gGI `       - sgG`+&gt;&gt;! tggv         &lt;OgRy .    ` JAgKe .      
      . w$gh;           ` [8gGI `         LgK4pppp4XggXn .      &lt;OgRy .    ` JAgKe .      
      . w$gh;           ` [8gGI `      . lH85xxxxxjZOBg( -      &lt;OgRy .    ` JAgKe .      
      . w$gh;           ` [8gGI `      - Fgk _"''', Cgg1"       &lt;OgRy .    ` JAgKe .      
      . w$gh;           ` [8gGI `       :3gO `      x$ggq .     &lt;OgRy .    ` JAgKe .      
      . q$gh&gt;           ` n8gb3 `     . 2gEc `    `  vggl -     +OgDa .    ` 7AgKo .      
      : E$gP=`-        .: 1mgOC :.   ., lgY :      : cggZ_ '  ' &gt;pgRx :    ' )AgA[ "      
        bggb/             4ggm]        LqgB          [R$g=      LKgg9        38gD]        
     r*2ZgBg)Z#`       hnmSgBgaHnq ` 2nIggnY'"     4}&gt;gBgb 8 -gZ)gBgZEc&lt;  '.YigBgla&lt;_     
   . jgggD#Dgg[_     ` :gggD#Dggg8 ' OggRDggD/ ` ` ,ggD#DRg&lt; ;YggD#Dggg[  +wggD#Dgggv `   
     =yj]222ya[`       f]j]222]j]{   F]y2Eayx'     f]y222]2} .nay222]jj^  -oya222ayx"
</pre>

# Go Image to ASCII Server

![Demo GIF](./assets/demo.gif)

## Table of content
- [Introduction](#introduction)
- [Two ways to run](#two-ways-to-run)
    - [Defalt](#default)
    - [Docker](#docker)

## Introduction
Upload an image, and get ASCII version of it using the simplest UI.

## There are two ways to run this
### Default
#### Prerequisites
* Go >= 1.22.0
* Gin >= 1.9.1
* VIPS >= 8.15.1
* pkg-config >= 0.29.2

The MacOS installation through [brew](https://brew.sh/):

```sh
make install
```

#### Running the server
```sh
# Development
make run-dev

# Production
make run-prod
```

### Docker
#### Prerequisites
* Docker >= 25.0.3

```sh
# Build the whole project.
make build

# Run the server.
make dock-run

# Stop the server.
make dock-stop

# Clean up (removes container and image)
make dock-clean
```
