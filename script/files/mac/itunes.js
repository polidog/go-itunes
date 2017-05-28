var app = Application('iTunes')
var event = Application('System Events')


function getCurrentTrack() {
    var track = app.currentTrack()
    return {
        name: track.name(),
        artist: track.artist(),
        album: track.album(),
        category: track.category(),
        time: track.time()
    }
}

function pausePlaying() {
    var result = {'status': true}
    if (isPlaying()) {
        app.pause();
        result['status'] = false;
    }
    result['status'] = true;
    result['player_state'] = app.playerState()
    return result;
}

function startPlaying() {
    if (!isPlaying()) {
        app.play();
        return {
            'status': true,
            'track': getCurrentTrack(),
            'player_state': app.playerState(),
        }
    }
    return {
        'status': false,
        'player_state': app.playerState(),
    };
}

function stopPlaying() {
    if (isPlaying()) {
        app.stop();
        return {
            'status': true,
            'player_state': app.playerState(),
        }
    }
    return {
        'status': false,
        'player_state': app.playerState(),
    }
}

function playNextTrack() {
    app.nextTrack();
    return {
        'status': true,
        'track': getCurrentTrack(),
        'player_state': app.playerState(),
    }
}

function playPreviousTrack() {
    app.previousTrack();
    return {
        'status': true,
        'track': getCurrentTrack(),
        'player_state': app.playerState(),
    }
}

function isPlaying() {
    return app.playerState() === 'playing'
}


function run(argv) {
    switch(argv[0]) {
        case "track":
            return JSON.stringify({
                'status': true,
                'track': getCurrentTrack()
            });
            break;
        case "pause":
            return JSON.stringify(pausePlaying())
            break;
        case "play":
            return JSON.stringify(startPlaying())
            break;
        case "stop":
            return JSON.stringify(stopPlaying())
            break;
        case "next":
            return JSON.stringify(playNextTrack())
            break;
        case "prev":
            return JSON.stringify(playPreviousTrack())
            break;
        default:
            return JSON.stringify({
                'status': false,
                'message': 'Unsupported command'
            })
            break;
    }
}
