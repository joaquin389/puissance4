class P4 {
    constructor(selector) {
        this.Col = 7;
        this.Lgn = 6;
        this.selector = selector;

        this.drawGame();
    }

    drawGame() {
        const $game = $(this.selector);

        for(let Lgn = 0; Lgn < this.Lgn; Lgn++) {
            const $row = $('<div>').addClass('row');
            for(let Col = 0; Col < this.Col; Col++) {
                const $col = $('<div>').addClass('cell').attr('data-col', Col).attr('data-lgn', Lgn);
                $row.append($col);
            }
            $game.append($row);
        }
    }
}