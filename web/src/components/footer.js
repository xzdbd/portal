import React from 'react'
import PropTypes from 'prop-types'
import { withStyles } from 'material-ui/styles'
import Typography from 'material-ui/Typography'
import Grid from 'material-ui/Grid'

const styles = {
  root: {
    flexGrow: 1,
    width: '100%',
    paddingTop: 40
  },
  footerGrid: {
    borderTop: '4px solid #4179F7',
    backgroundColor: '#F2F6FF',
    minHeight: 110
  },
  image: {
    maxWidth: 200
  }
}

function Footer (props) {
  const { classes } = props
  return (
    <footer className={classes.root}>
      <Grid container spacing={0} justify='center'>
        <Grid item xs={12} className={classes.image}>
          <img src='https://slidesup.com/static/media/mascot.79381efd.svg' />
        </Grid>
        <Grid item xs={12} className={classes.footerGrid}>
          <Typography variant='subheading' color='inherit' align='center'>
                        Made with ❤️ in HZ
          </Typography>
          <Typography variant='caption' color='inherit' align='center'>
                        Copyright © 2018 xzdbd
          </Typography>
        </Grid>
      </Grid>
    </footer>
  )
}

Footer.propTypes = {
  classes: PropTypes.object.isRequired
}

export default withStyles(styles)(Footer)
