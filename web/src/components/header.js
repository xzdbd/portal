import React from 'react'
import PropTypes from 'prop-types'
import { withStyles } from 'material-ui/styles'
import AppBar from 'material-ui/AppBar'
import Toolbar from 'material-ui/Toolbar'
import Typography from 'material-ui/Typography'
import Grid from 'material-ui/Grid'
import Button from 'material-ui/Button'

const styles = {
  root: {
    flexGrow: 1,
    backgroundColor: '#F2F6FF'
  },
  flex: {
    flex: 1
  },
  appBar: {
    background: '#F2F6FF'
  }
}

function HeaderBar (props) {
  const { classes } = props
  return (
    <div className={classes.root}>
      <Grid container spacing={0} justify='center'>
        <Grid item xs={8} >
          <AppBar classes={{ colorDefault: props.classes.appBar }} position='static' color='default' elevation={0}>
            <Toolbar>
              <Typography variant='title' color='inherit' className={classes.flex}>
                <img src='images/logo.svg' />
              </Typography>
              <Button color='primary'><b>Login</b></Button>
              <Button color='primary' variant='raised'><b>Document</b></Button>
            </Toolbar>
          </AppBar>
        </Grid>
      </Grid>
    </div>
  )
}

HeaderBar.propTypes = {
  classes: PropTypes.object.isRequired
}

export default withStyles(styles)(HeaderBar)
