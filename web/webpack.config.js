module.exports = {
    mode: 'development',
    entry: ['babel-polyfill', './src/entry.js'],
    output: {
        path: __dirname + '/',
        filename: 'index.js',
    },
    devServer: {
        inline: true,
        contentBase: './dist',
        port: 3000
    },
    module: {
        rules: [
            {
                test: /\.js$/,
                exclude: /node_module/,
                use: {
                    loader: 'babel-loader',
                    options: {
                        presets: ['es2015', 'react', 'stage-0']
                    }
                }        
            }
        ]
    },
};