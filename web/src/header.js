import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import AppBar from 'material-ui/AppBar';
import Toolbar from 'material-ui/Toolbar';
import Typography from 'material-ui/Typography';
import Grid from 'material-ui/Grid';

const styles = {
  root: {
    flexGrow: 1,
  },
  appBar: {
    background: '#fafafa',
  },
};

function HeaderBar(props) {
  const { classes } = props;
  return (
    <div className={classes.root}>
        <Grid container spacing={0} justify='center'>
            <Grid item xs={6} >
                <AppBar classes={{colorDefault: props.classes.appBar}} position="static" color="default" elevation={0}>
                    <Toolbar>
                    <Typography variant="title" color="inherit">
                        Title
                    </Typography>
                    </Toolbar>
                </AppBar>
            </Grid>
        </Grid>
    </div>
  );
}

HeaderBar.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(HeaderBar);