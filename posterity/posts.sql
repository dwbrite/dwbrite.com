
--drop table posts cascade;

INSERT INTO posts(title, post_date, body)
VALUES('The start of something',
       '2015-06-11',
       '<p>Last month I started re-playing Pokemon Fire Red on the VBA-M emulator. ' ||
       'I went into the game with the goal of catching as many Pokemon as possible. ' ||
       'After about two weeks I managed to catch 124 of the 151 in the game. ' ||
       'That magic number also happened to be the maximum number of Pokemon you could get in either of the games.</p>' ||
       '' ||
       '<p>I still wanted to catch all the Pokemon though. ' ||
       'I’d never really be happy with it until I had all of them. ' ||
       'So I moved my save to Leaf Green and caught all the LG exclusives. ' ||
       'This brought me to a total of 138 caught, 149 seen. ' ||
       'I should mention that I would also only capture Female Pokemon ' ||
       'and that I have one of every Pokemon in my box in numeric order, ' ||
       'but that’s besides the point.</p>' ||
       '' ||
       '<p>The only Pokemon I couldn’t get were the 4 trade-evolvers, the two other starters, ' ||
       'the other fossil Pokemon (Omanyte), and the legendary Mew. ' ||
       'Omastar and Mew were the only Pokemon I hadn’t seen.</p>' ||
       '' ||
       '<p>And I wanted them.</p>' ||
       '' ||
       '<p>What really sucks though is that trading on VBA-M is and has been broken forever on FR/LG. ' ||
       'I thought: “Hey, I program, I should help out with this and figure out the glitch!” so I did just that. ' ||
       'Kind of… I opened up the source code in Visual Studio and got it to compile right - Success! ' ||
       'After a little while scouring the source I realized that I really didn’t know C++ well at all.</p>' ||
       '' ||
       '<p>I gave up on that and decided I would make my own Pokemon game. ' ||
       'That it should run on all systems, be able to use a controller, have online trades, ' ||
       'have all the regions, and most importantly that you could catch all the Pokemon.</p>' ||
       '' ||
       '<p>Five days later and here I am starting this dev blog.</p>' ||
       '' ||
       '<p>But I’ve never really finished a project this big before, ' ||
       'so I don’t want this to get popular until I finish. ' ||
       'I don’t want to disappoint a lot of people. ' ||
       'If you find this I beg you not to get too excited until I’ve made considerable progress.</p>' ||
       '' ||
       '<p>I do want to finish this though, ' ||
       'so if anyone stumbles upon this and notices I’m not updating the blog or working on the game, ' ||
       'please spam my email or telephone or anything you can.</p>' ||
       '' ||
       '<p>Email: dwbrite@gmail</p>' ||
       '' ||
       '<p>Thanks~</p>');


INSERT INTO posts(title, post_date, body)
VALUES('Pocket Progress 1',
       '2015-06-22',
       '<noscript class="media-noscript">' ||
       '  <img src="/resources/media/pocket-progress/first-screenshot.png" alt="image"/>' ||
       '</noscript>' ||
       '' ||
       '<p>It may not look like much but this is really exciting progress for me and this project. ' ||
       'This is a Pokemon recreation I’m making with Slick2D/Java that I’ve been working on for the past two weeks. ' ||
       'There’s no collisions or animations quite yet but those should be done within the next two days.</p>' ||
       '' ||
       '<p>This game is a recreation of FR/LG that’s expanded to ' ||
       'several other regions in a single story line including Hoenn, Johto, and Orre. ' ||
       'The main goal is to have a “pure” Pokemon game in which all Pokemon can be found ' ||
       'and there are no weird Dex restrictions ' ||
       '(For example, not being able to evolve a Crobat in FR/LG before the national dex, ' ||
       'or having game specific Pokemon like Vulpix or Bellsprout).</p>' ||
       '' ||
       '<p>The biggest problem is leveling throughout regions. ' ||
       'I think the best solution is to have a level lock until you beat certain gyms - ' ||
       'similar to how traded Pokemon of a high level refuse to listen.</p>' ||
       '' ||
       '<p>If you have any ideas send me an email and I’ll think about it. ' ||
       'Remember to check my live update page for more information on what exactly is being done!</p>');


INSERT INTO posts(title, post_date, body)
VALUES('Pocket Progress 2',
       '2015-06-27',
       '<noscript class="media-noscript">' ||
       '  <img src="/resources/media/pocket-progress/collision-and-animation.gif" alt="image"/>' ||
       '</noscript>' ||
       '' ||
       '<p>Collision and animations are looking good right now. ' ||
       'Here’s a gif of where I’m at. ' ||
       'It took me all day to figure out what the hell I was doing wrong with interrupting walking into walls. </p>' ||
       '' ||
       '<p>Anyway, I also overhauled the movement and controls I’d put in yesterday ' ||
       'because I realized exactly how the real games handled it.</p>' ||
       '' ||
       '<p>You can see travelling through areas right in this, which is nice, ' ||
       'and you can also see that I haven’t programmed in collision for cliffs yet. :x</p>' ||
       '' ||
       '<p>I had to keep the gif at 15FPS to make sure it was under 2MB while not looking being mangled by compression, ' ||
       'but the game does in fact run at 60FPS, just like a real Gameboy game.</p>' ||
       '' ||
       '<p>That’s pretty much it for now. ' ||
       'Make sure to check the Live Updates page on my blog for live updates of what I’ve done or am doing. Bye!</p>');


INSERT INTO posts(title, post_date, body)
VALUES('Pocket Progress 3',
       '2015-07-06',
       '<noscript class="media-noscript">' ||
       '  <img src="/resources/media/pocket-progress/collision-and-animation2.gif" alt="image"/>' ||
       '</noscript>' ||
       '' ||
       '<p>Look at how far we’ve come! It’s not really that much but it’s exciting to me. ' ||
       'Here you can see animations for running, biking, and cliff jumping in action! ' ||
       'Cliff jumping still needs the offset for the sprite, the shadow, ' ||
       'and the dust particles when you land, but the movement and animation frames are all in place!</p>' ||
       '' ||
       '<p>This gif is a bit lower quality than the last one, even though it’s still at 15FPS, ' ||
       'but it gets the point across. Anyway, that’s it for now. </p>' ||
       '' ||
       '<p><s>Make sure to follow if you’re interested in this project. ' ||
       'If you want to know what’s being done every day follow my ' ||
       '<a href="http://dwbriteupdates.tumblr.com/" target="_blank">daily updates</a> page.</s></p>' ||
       '' ||
       '<p>It isn’t shown in the gif, but you can no longer phase through the cliffs going upwards.</p>');


INSERT INTO posts(title, post_date, body)
VALUES('I’m back',
       '2015-07-25',
       '<p>' ||
       'Probably. I forget what happened, but stuff happened and my sleep schedule got messed up, ' ||
       'and I just fixed it the other day and that’s my excuse, ' ||
       'but I’m here now and I can finally sleep at a decent time.</p>' ||
       '' ||
       '<p>Anyway&hellip; I’ve finally got grass working beautifully, ' ||
       'and I actually discovered exactly how grass works through a funky glitch in FR/LG.</p>' ||
       '' ||
       '<p>Here’s how it works: Use cut in the north-most grass in Route 1, ' ||
       'take one step into Viridian city, then go back down to the grass. ' ||
       'When you step on the empty tiles, ' ||
       'the grass particle effects show and you still have a chance of wild encounters. ' ||
       'I figured out that rather than having three separate particles ' ||
       'for each “piece” of the animation ' ||
       '(the stepped on floor, the fluttering grass/leaves, and the part that stays on above your character), ' ||
       'it’s all in one particle, that looks like this:  </p>' ||
       '' ||
       '<noscript class="media-noscript">' ||
       '  <img alt="image" src="/resources/media/pocket-progress/grass-rustle.png"/>' ||
       '</noscript>' ||
       '' ||
       '<p>and changes the depth during the animation. ' ||
       'There’s just one frame in which the second image is behind the character which led me to the discovery.</p>' ||
       '' ||
       '<p>The glitch isn’t listed anywhere I could find easily, ' ||
       'which reminds me of another visual glitch in Vermilion City, which is less interesting, ' ||
       'but showcases a glaring flaw with the layer system used in the generation 3 games.</p>' ||
       '' ||
       '<noscript class="media-noscript">' ||
       '  <img alt="image" src="/resources/media/pocket-progress/frlg-layer-glitch.png"/>' ||
       '</noscript>' ||
       '' ||
       '<p>And lastly, here’s a comparison between what my game looks like now, versus what the real game looks like. ' ||
       'Pretty close, right? Forgive the input lag on the left screen (The real game).</p>' ||
       '' ||
       '<noscript class="media-noscript">' ||
       '  <img alt="image" src="/resources/media/pocket-progress/cmp-frlg.gif"/>' ||
       '</noscript>');


INSERT INTO posts(title, post_date, body)
VALUES('Updates',
       '2015-08-01',
       '<ul>' ||
       '<li>WINDOWS 10 BABY</li>' ||
       '<li>Web server is back up: ' ||
       '<a href="http://dwbrite.com" target="_blank">http://dwbrite.com</a> or ' ||
       '<a href="http://www.dwbrite.com" target="_blank">http://www.dwbrite.com</a>. ' ||
       'I have an SSL cert, so you can try to go to https, ' ||
       'but until I get a more legitimate cert I’m going to have it block the rest of the site.</li>' ||
       '<li>You can also try emailing me: devin@dwbrite.com, ' ||
       'but because I don’t have a static IP I can’t guarantee it’s not blacklisted, ' ||
       'which would mean you don’t get my replies.</li>' ||
       '<li>Made a desktop from a bunch of spare parts, running Zorin OS 9 on an Athlon 64. ' ||
       'Right now I’m using it as a media center. ' ||
       'I might get XMBC on it.</li><li>Plugged in flashy lights.</li>' ||
       '</ul>' ||
       '' ||
       '<p>____________________</p>' ||
       '' ||
       '<p>TODO:</p>' ||
       '<ul>' ||
       '<li><s>Paint wall(s)</s></li>' ||
       '<li><s>Rearrange room.</s></li>' ||
       '<li>Listen to more hip hop and EDM.<br/></li>' ||
       '</ul>');


INSERT INTO posts(title, post_date, body)
VALUES('GifBackgroundifier!',
       '2015-08-06',
       '<noscript class="media-noscript">' ||
       '  <img src="/resources/media/gifbackgroundifier.gif" alt="image"/>' ||
       '</noscript>' ||
       '' ||
       '<p>Finished my “GifBackgroundifier”! Set an animated .gif as your background on a Windows computer! ' ||
       'I think it’s pretty rad, and if you’d like to try it just ask.</p>' ||
       '' ||
       '<hr>' ||
       '' ||
       '<p>Download link is here: ' ||
       '<a href="https://www.dropbox.com/s/88sr7rt5wk39of0/GifBackgroundifier.jar?dl=0" target="_blank">' ||
       'https://www.dropbox.com/s/88sr7rt5wk39of0/GifBackgroundifier.jar?dl=0' ||
       '</a></p>' ||
       '' ||
       '<p>If your GIF doesn’t work, try opening it in GIMP and exporting it to a new file with the highest quality. ' ||
       'This is just the first version so I haven’t made an error popup, ' ||
       'but you can tell if your GIF doesn’t work by if the loading bar moves or not.</p>' ||
       '' ||
       '<p>Windows only, sorry.<br/></p>' ||
       '' ||
       '<p>Feel free to send me ideas, suggestions, or complaints!<br/></p>' ||
       '' ||
       '<p>Warning: This program uses a decent amount of processing power, about 12% on my laptop.<br/></p>');


INSERT INTO posts(title, post_date, body)
VALUES('Research and Whiteboards?',
       '2015-11-07',
       '<p>I’ve been yelling at myself to write something for a while now, ' ||
       'but I’ve been putting it off because I haven’t really figured out what to write. ' ||
       'So I guess I’ll just wing it and this will be whatever it is.</p>' ||
       '' ||
       '<p>Over the past however-long-it’s-been-since-my-last-post I haven’t been doing nothing. ' ||
       'I mean I’ve mostly been doing nothing, but I haven’t been doing completely nothing. ' ||
       'When I reached the point in my Pokemon remake where I had to ' ||
       'gather all of the data for the Pokemon I got lazy. ' ||
       'I really don’t like tedious work, so I kind of ignored it for a while. ' ||
       'In the mean time I did a few other things to kill time.</p>' ||
       '' ||
       '<p>The very first thing I made since I’ve been gone is a glowey-whiteboard. ' ||
       'Here’s the final result:</p>' ||
       '' ||
       '<noscript class="media-noscript">' ||
       '  <img alt="image" src="/resources/media/glowboard-0.jpg"/>' ||
       '</noscript>' ||
       '' ||
       '<noscript class="media-noscript">' ||
       '  <img alt="image" src="/resources/media/glowboard-1.jpg"/>' ||
       '</noscript>' ||
       '' ||
       '<p> The whole build needs a post of its own so I’ll try to do that within the next week.</p>' ||
       '' ||
       '<p>After that I got back on track a bit. ' ||
       'I downloaded a bunch of sprites from ' ||
       '<a href="http://veekun.com/dex/downloads" target="_blank">Veekun’s collection of downloads</a>. ' ||
       'I also wanted to get the animated sprites from Black/White, ' ||
       'so I used HTTP GET requests to gather those from Pokemondb. ' ||
       'It took a while and it ended up being something like 65MB total. ' ||
       'Not terrible, but I can understand why Veekun wouldn’t want that on his server.</p>' ||
       '' ||
       '<p>At around the same time I finally decided to get all of the Pokemon data up until gen 3. ' ||
       'Basically I made a list of all the data needed for each Pokemon ' ||
       '(Index #s, names, gender rates, color, egg types, etc.), and put it all into one convenient spreadsheet. ' ||
       'I have some different pages for certain data like moves, evolutions, and forms - ' ||
       'but I haven’t filled out forms and moves yet. ' ||
       'I still have to decide how I’m going to program those in first. ' ||
       'For forms/transformation I’ve been tossing around this data structure that seems to fit how GameFreak does it. ' ||
       'There are three form interfaces: “In-battle Transforming” (Castform, Cherrim, Giratina), ' ||
       '“Evolution Transforming” (Deoxys, Burmy, Rotom), ' ||
       '“Non-Transforming” (Unown, Shellos). ' ||
       'That spreadsheet can be found ' ||
       '<a href="https://docs.google.com/spreadsheets/d/1R9KzpWbyOp9GM-laW3as4nkQBD2jWWaAT8XgrpFzJN0/edit?usp=sharing" target="_blank">here</a>. ' ||
       'There’s probably some worksheet pages that I used just to paste some regex data, so you can ignore those.</p>' ||
       '' ||
       '<p>I also added the particles for ledge-jumping and running through grass, ' ||
       'and I’m in the process of adding battles right now, ' ||
       'I just need to figure out how I’m going to implement it. ' ||
       'Here’s a beautifully looped gif of the game in its current form.</p>' ||
       '' ||
       '<noscript class="media-noscript">' ||
       '  <img alt="image" src="/resources/media/pocket-progress/ledge-hopping.gif"/>' ||
       '</noscript>' ||
       '' ||
       '<p>I also painted and reorganized my room, and added some neat lights, ' ||
       'which I’ll show sometime in the future whenever I get a chance.</p>' ||
       '' ||
       '<p>That’s about it for the past. Now onto the future. ' ||
       'My current non-Pokemon projects are making a smart mirror ' ||
       'from an old laptop screen and some minicomputer, and making a custom aux cable. ' ||
       'The smart mirror is <a href="http://imgur.com/gallery/dO8Yl" target="_blank">fairly straightforward</a>, ' ||
       'but the aux cable is where it gets interesting.</p>' ||
       '' ||
       '<p>I’ve always been frustrated by headphones. ' ||
       'Especially expensive ones with hardware you can’t replace. ' ||
       'The most notably example being the cord, as simple as it is. ' ||
       'They always seem to break, usually right at the base of the jack, ' ||
       'completely ruining a nice pair of headphones. ' ||
       'The solution is usually to cut it off and solder on a new jack, ' ||
       'which shortens the cord and usually comes out pretty ugly. ' ||
       'Another thing that bothers me is the length of the wire. ' ||
       'They’re either too short for anything but tying an iPod to your arm, ' ||
       'or so long that they dangle down to your knees. ' ||
       'I might be exaggerating a bit, but it begs the question: ' ||
       'why aren’t headphone wires made to be retractable? Obviously not buds though.</p>' ||
       '' ||
       '<p>So I’ve decided I’m going to ' ||
       'a.) mount a female TRS jack where the wire splits, and ' ||
       'b.) make a custom TRS/aux cable. ' ||
       'As a personal choice I’d like to have it with a right-angle jack on one end, ' ||
       'a straight jack on the other, and a paracord sheath. ' ||
       'I also want it to have a retractable coil. ' ||
       'This can all be done by <a href="http://pexonpcs.co.uk/" target="_blank">Pexon</a>, ' ||
       'but it’s cheaper to do on my own and should offer some valuable experience.</p>' ||
       '' ||
       '<p>The idea is to get some magnet wire, twist those together, ' ||
       'wrap that in aluminum coated Mylar to prevent interference, ' ||
       'cover that in a polyvinyl sleeve, then cover that in a Nylon paracord sheath, ' ||
       'wrap the cable around a metal rod and bake it hot enough for heat memory, ' ||
       'then twist it in the opposite direction of the coil to make it snappy, and put on the jacks. ' ||
       'Pretty easy. The cable should be 10′ long total, and should compress to about 3′. ' ||
       'Obviously I haven’t included all the details, but I’ll give the nitty-gritty when I make a full post on this. ' ||
       'Overall it will cost about $50, and I’ll be able to make two cables, ' ||
       'with each new one after that costing &lt;$20. ' ||
       'All that’s needed is more PVC, Mylar, paracord, and jacks. ' ||
       'Much cheaper than the $70+ for a similar cord from Pexon.</p>' ||
       '' ||
       '<p>I guess that’s about it. I already know what I’ll be working on when I finish all of everything, ' ||
       'but I’m going to keep that on the hush hush until I at least finish this Pokemon game. ' ||
       'That’s it for now so until the next time I decide to stop being a lazy bastard, goodbye and goodnight.</p>');


INSERT INTO posts(title, post_date, body)
VALUES('Quick Transition',
       '2015-12-30',
       '<p>A few days ago I transferred my domains from GoDaddy to Google Domains, ' ||
       'and while the process isn’t quite complete yet, it should be done pretty soon. ' ||
       'With this, I’ll also be changing the format of the website a little bit. ' ||
       'The daily updates page hasn’t really been working out, ' ||
       'so I’ll be converting that into a personal blog. ' ||
       'I often find that I just want to say something that’s not really important, ' ||
       'and I simply can’t find a reason to clutter my dev log with it.</p>' ||
       '' ||
       '<p><s>I’ll be separating the site into separate sub-domains for each part. ' ||
       'The dev log will become dev.dwbrite.com, the blog will become blog.dwbrite.com, ' ||
       'and the about me will become simply dwbrite.com. ' ||
       'Then I think I’ll just add a portfolio and everything should be set. ' ||
       'Things might get ugly for a minute but just pretend nothing happened. :)</s></p>' ||
       '' ||
       '<p>For actual projects/development, I’m finally starting the headphone/audio cable! ' ||
       'All the parts should be here by the 14th at the latest so if everything goes as planned ' ||
       'I’ll get a post up shortly after.</p><p>On a similar note, ' ||
       'I haven’t made a post regarding the creation of the light board. ' ||
       'For that I’m sorry, but I do have an explanation. ' ||
       'I <b><i>really</i></b> don’t want to take it all down, and I haven’t had a reason to make another one. ' ||
       'They’re pretty darn expensive after all!</p>' ||
       '' ||
       '<p>I guess that’s all I have to say for now, so Merry Christmas and Happy New Year! ' ||
       'I’ll be back in a few weeks!</p>');


INSERT INTO posts(title, post_date, body)
VALUES('November Has Come',
       '2017-11-19',
       '<p>Wow. I can’t believe it’s been over two years since my last blog post. ' ||
       'I’ve learned a lot since the start of this blog. Well, here I am. ' ||
       'I read a bunch of books on programming. I got my diploma and started working on a degree. ' ||
       'Man, how time flies.</p>' ||
       '' ||
       '<p>I’m trying to be more conscious of the world around me. ' ||
       'I’m using Linux whenever possible, but some classes <em>require</em> otherwise. ' ||
       'Luckily spinning up a VM is never too hard. ' ||
       'I’ve tested out every desktop environment under the sea before finally settling on Gnome. ' ||
       'I still can’t find a suitable terminal or file manager. ' ||
       'These are two things that macOS seem to get right, but which no-one else can.</p>' ||
       '' ||
       '<p>Ligatures! I love them. The moz://a foundation is a developer’s best friend. ' ||
       'FiraMono/FiraCode. Rust. Firefox nightly. Just&hellip; Everything. ' ||
       'They’re so great right now! I’m glad all their efforts are catching up to them :)</p>' ||
       '' ||
       '<p>I’ve also learned a bit of OpenGL. ' ||
       'I bought a domain and made <a href="https://ajitvpai.com" target="_blank">a website</a> ' ||
       'which blew up on reddit for a minute. I started learning Rust and Kotlin.</p>' ||
       '' ||
       '<p>The Pokemon project is still ongoing! ' ||
       'It’s currently in a private github repository, ' ||
       'and I think I’ll release it when the code is pretty enough. ' ||
       'I’m refactoring the old code, updating it with new knowledge, ' ||
       'and replacing Java with Kotlin as I go. ' ||
       'Planning for battles - the largest part of development - is nearly complete. ' ||
       'Everything else should move along pretty quickly, ' ||
       'and progress should <em>look</em> much faster when that happens.</p>' ||
       '' ||
       '<p>I’m sure I’m forgetting some things, but this will do for now.</p>');
