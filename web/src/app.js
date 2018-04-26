import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import Paper from 'material-ui/Paper';
import Grid from 'material-ui/Grid';
import HeaderBar from './header'

const styles = theme => ({
    root: {
        flexGrow: 1,
    },
    paper: {
        padding: theme.spacing.unit * 2,
        textAlign: 'center',
        color: theme.palette.text.secondary,
    },
});

function FullWidthGrid(props) {
    const { classes } = props;

    return (
        <div className={classes.root}>
            <Grid container spacing={0}>
                <Grid item xs={12} >
                    <HeaderBar />
                </Grid>
                <Grid item xs={12}>
                    <Paper className={classes.paper}>search</Paper>
                </Grid>
                <Grid item xs={12}>
                    <Paper className={classes.paper}>items</Paper>
                </Grid>
                <Grid item xs={12}>
                    <Paper className={classes.paper}>footer</Paper>
                </Grid>
                
            </Grid>
        </div>
    );
}

FullWidthGrid.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(FullWidthGrid);