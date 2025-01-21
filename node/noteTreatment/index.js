// test node
// var http = require('http');

// http.createServer(function (req, res) {
//     res.writeHead(200, {'content-type' : 'text/html'});
//     res.end('Hello World');
// }).listen(8080);

// file system for file manipulation
var fs = require('fs');
// read line // std::cin or Console.ReadLine();
const readline = require('readline');

// standard I/O stram
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

function AddNote()
{
    // add note to file
    const Note = (note) => {
        fs.writeFile('note.txt', note, (err) => {
            if (err) throw err;
        })
    }

    // Add user input to note.txt
    rl.question('Write a note: ', (note) => {
        rl.close();
        Note(note);
    })
}

function ReadNotes()
{
    // Read file
    fs.readFile('./note.txt', 'utf8', (err, data) => {
        if (err) throw err;
        console.log(data);
    })
}

function DeleteNote()
{
    // delete the wanted line in note.txt
    const DeleteLine = (data, line) => {
        return data
            .split('\n')
            .filter((_, idx) => idx !== parseInt(line, 10))
            .join('\n');
    }

    // take user input so that the note can be deleted
    rl.question('Write the number of the line you want to delete: ', (line) => {
        fs.readFile('./note.txt', 'utf8', (err, data) => {
            if (err) throw err;

            const UpdatedNote = DeleteLine(data, line);

            fs.writeFile('./note.txt', UpdatedNote, 'utf8', (err) => {
                if (err) throw err;

                console.log('The note has been deleted!');
                rl.close();
            })
        })
    })
}

NoteAppMenu()
{
    console.log("[1] Add note\n[2] Read notes\n[3] Delete note");
    rl.question(">", (input) => {
        switch (input)
        {
            case 1:
            AddNote();
            NoteAppMenu();
            break;
            case 2:
            ReadNotes();
            NoteAppMenu();
            break;
            case 3:
            DeleteNote();
            NoteAppMenu();
            break;
            default:
            console.log("Invalid input");
            NoteAppMenu();
            break;
        }
    })
}

NoteAppMenu();
