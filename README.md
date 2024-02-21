
# Dragon Mail

Secure E-Mail client with a breath of destruction.


# Libraries

- https://pkg.go.dev/github.com/emersion/go-imap/v2/imapclient#Client

# (Planned) Features

- No insecure connections, only TLS
- IMAP Support
- PGP Integration
- Multiple Identities
- Sender Pinning (IP to Contact Book matching)
- Plain Text Only, no HTML
- Markdown Mode (for HTML responses)

- Import of E-Mails
- Thunderbird (Maildir/mbox)
- Outlook (pst)

- E-Mail Filter
-> Move email with keyword xyz to folder xyz
-> Move email from @domain to folder xyz


# Profile Folder Structure

- /account.json
- /account/domain/sender/2021-12-31 - 123834854858.eml

# Ascii Art
Art by Shanaka Dias
              __
          _.-'.-'-.__
       .-'.       '-.'-._ __.--._
-..'\,-,/..-  _         .'   \   '----._
 ). /_ _\' ( ' '.         '-  '/'-----._'-.__
 '..'     '-r   _      .-.       '-._ \
 '.\. Y .).'       ( .'  .      .\          '\'.
 .-')'|'/'-.        \)    )      '',_      _.c_.\
   .<, ,>.          |   _/\        . ',   :   : \\
  .' \_/ '.        /  .'   |          '.     .'  \)
                  / .-'    '-.        : \   _;   ||
                 / /    _     \_      '.'\ ' /   ||
                /.'   .'        \_      .|   \   \|
               / /   /      __.---'      '._  ;  ||
              /.'  _:-.____< ,_           '.\ \  ||
             // .-'     '-.__  '-'-\_      '.\/_ \|
            ( };====.===-==='        '.    .  \\: \
             \\ '._        /          :   ,'   )\_ \
              \\   '------/            \ .    /   )/
               \|        _|             )Y    |   /
                \\      \             .','   /  ,/
                 \\    _/            /     _/
                  \\   \           .'    .'
                   '| '1          /    .'
                     '. \        |:    /
                       \ |       /', .'
                        \(      ( ;z'
                         \:      \ '(_
                          \_,     '._ '-.___
                snd                   '-' -.\