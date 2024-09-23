const std = @import("std");

fn nextLine(reader: anytype, buffer: []u8) !?[]const u8 {
    const line = (try reader.readUntilDelimiterOrEof(
        buffer,
        '\n',
    )) orelse return null;
    // trim annoying windows-only carriage return character
    if (@import("builtin").os.tag == .windows) {
        return std.mem.trimRight(u8, line, "\r");
    } else {
        return line;
    }
}

fn metaCmd(input: []const u8) !void {
    if (std.mem.eql(u8, input, ".quit")) {
        std.process.exit(0);
        unreachable;
    } else if (std.mem.eql(u8, input, ".info")) {
        const stdout = std.io.getStdOut();
        try stdout.writeAll(
            \\   •      ┓
            \\ ┏┓┓┏┏┓┏┏┓┃
            \\ ┣┛┗┗┗┛┛┗┫┗
            \\ ┛       ┗
            \\ a tiny sqlite clone
            \\ 
            \\
        );
    } else {
        return error.UnknownCommand;
    }
}

pub fn main() !void {
    const stdout = std.io.getStdOut();
    const stdin = std.io.getStdIn();

    try stdout.writeAll("repl started. type `.quit` to exit\n");

    // main repl loop
    while (true) {
        try stdout.writeAll(">>> ");
        var buffer: [100]u8 = undefined;
        const input = (try nextLine(stdin.reader(), &buffer)).?;

        // meta commands
        if (input[0] == 46) {
            metaCmd(input) catch try stdout.writeAll("Unrecognised Command\n");
        }

        // try stdout.writer().print(
        //     "{s}\n",
        //     .{input},
        // );
    }
}
